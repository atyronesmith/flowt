package schema

import (
	"fmt"
	"go/format"
	"math"
	"reflect"
	"runtime"
	"sort"
	"strings"

	types "github.com/atyronesmith/flowt/pkg/ovsdbflow"
	"github.com/atyronesmith/flowt/pkg/utils"
)

type CType int

const (
	Atomic CType = iota
	Map
	Set
)

func mapAtomic(aType string, cType CType) (string, error) {
	var rType string

	switch aType {
	case "string":
		rType = "string"
	case "integer":
		rType = "int"
	case "real":
		rType = "float"
	case "boolean":
		rType = "bool"
	case "uuid":
		rType = "types.UUID"
		if cType == Atomic {
			cType = Set
		}
	default:
		return "", fmt.Errorf("unknown type: %s", aType)
	}

	switch cType {
	case Atomic:
		return rType, nil
	case Map:
		return fmt.Sprintf("types.OVSMap[%s]", rType), nil
	case Set:
		return fmt.Sprintf("types.OVSSet[%s]", rType), nil
	}
	return "", fmt.Errorf("unknown type: %s", aType)
}

func parseType(t interface{}) (string, error) {

	k := reflect.ValueOf(t).Kind()

	if k == reflect.String {
		// An atomic type, "integer", "real", "boolean", "string", or "uuid"
		return mapAtomic(t.(string), Atomic)
	} else if k == reflect.Map {
		tMap := t.(map[string]interface{})

		value, hasValue := tMap["value"]
		key := tMap["key"]

		if hasValue {
			if reflect.ValueOf(value).Kind() == reflect.String {
				if value.(string) != "string" {
					function, file, line, _ := runtime.Caller(1)
					fmt.Printf("%s, %s, %d\n", runtime.FuncForPC(function).Name(), file, line)
				}
				return mapAtomic(value.(string), Map)
			} else if reflect.ValueOf(value).Kind() == reflect.Map {
				vMap := value.(map[string]interface{})
				if vType, ok := vMap["type"]; ok {
					return mapAtomic(vType.(string), Map)
				} else {
					return "", fmt.Errorf("cannot parse type: %v", vMap)
				}
			}
		} else {
			// min is absent, 0, or 1
			min, hasMin := tMap["min"]
			var minInt int = 1
			if hasMin {
				minInt = int(min.(float64))
			}
			// max is absent, > 0, or "unlimited"
			max, hasMax := tMap["max"]
			var maxInt int = 1
			if hasMax {
				if reflect.ValueOf(max).Kind() == reflect.String {
					maxInt = math.MaxInt // Unlimited case
				} else {
					maxInt = int(max.(float64))
				}
			}
			// If min == 0 && max == 1, and key is not a map --> optional scalar of type key
			// if min == 0, and max == "unlimited", and key is not a map --> set of type key:type
			// if min == 0, and key is a map
			//    map->type == "uuid" then set of uuid(string)
			//    otherwise atomic(map->type)
			if minInt == 0 && maxInt == 1 {
				if reflect.ValueOf(key).Kind() == reflect.Map {
					km := key.(map[string]interface{})

					return mapAtomic(km["type"].(string), Atomic)
				} else {
					return mapAtomic(key.(string), Atomic)
				}
			} else if minInt == 1 && maxInt == 1 {
				if reflect.ValueOf(key).Kind() == reflect.String {
					return mapAtomic(key.(string), Atomic)
				} else {
					t := key.(map[string]interface{})

					return mapAtomic(t["type"].(string), Atomic)
				}
			} else if maxInt == math.MaxInt {
				if reflect.ValueOf(key).Kind() == reflect.String {
					return mapAtomic(key.(string), Set)
				} else {
					t := key.(map[string]interface{})

					return mapAtomic(t["type"].(string), Set)
				}
			} else {
				return "", fmt.Errorf("could not parse type, min: %d, max: %x", minInt, maxInt)
			}
		}

	}

	return "", fmt.Errorf("unable to parse type: %v", t)
}

func ParseSchema(tbl types.OVSdbSchema) error {

	var tblStructStr strings.Builder
	var dbStructStr strings.Builder

	switch tbl.Type {
	case types.NB:
		tblStructStr.WriteString("package nb\n\n")
	case types.SB:
		tblStructStr.WriteString("package sb\n\n")
	default:
		return fmt.Errorf("unknown database type: %s", tbl.Type)
	}

	tblStructStr.WriteString("import (\n\ttypes \"github.com/atyronesmith/flowt/pkg/ovsdbflow\")\n")

	dbStructStr.WriteString(fmt.Sprintf("type %s struct {\n", utils.SnakeToCamel( tbl.Type.String() )))
	dbStructStr.WriteString("Date types.Time `json:\"_date\"`\n")

	tblMap := tbl.Tables

	tableKeys := make([]string, 0, len(tblMap))
	for k := range tblMap {
		tableKeys = append(tableKeys, k)
	}

	sort.Strings(tableKeys)

	for _, tblName := range tableKeys {
		tblDefIntrf := tblMap[tblName]

		tblStructStr.WriteString(fmt.Sprintf("type %s struct {\n", utils.SnakeToCamel(tblName)))

		dbStructStr.WriteString(fmt.Sprintf("%s map[string]%s `json:\"%s\"`\n", utils.SnakeToCamel(tblName), utils.SnakeToCamel(tblName), tblName))

		tableDef := tblDefIntrf.(map[string]interface{})
		if colIntrf, ok := tableDef["columns"]; ok {
			colMap := colIntrf.(map[string]interface{})
			colKeys := make([]string, 0, len(colMap))
			for k := range colMap {
				colKeys = append(colKeys, k)
			}
			sort.Strings(colKeys)
			for _, colName := range colKeys {
				colDefIntrf := colMap[colName]
				colDef := colDefIntrf.(map[string]interface{})
				if typeIntrf := colDef["type"]; ok {
					if pt, err := parseType(typeIntrf); err == nil {
						tblStructStr.WriteString(fmt.Sprintf("%s %s `json:\"%s\"`\n", utils.SnakeToCamel(colName), pt, colName))
					}
				}
			}
		}
		tblStructStr.WriteString("}")
	}
	dbStructStr.WriteString("}")

	formatted, err := format.Source([]byte(tblStructStr.String()))
	if err != nil {
		return fmt.Errorf("unable to gofmt")
	}
	fmt.Printf("%s", string(formatted))

	formatted, err = format.Source([]byte(dbStructStr.String()))
	if err != nil {
		return fmt.Errorf("unable to gofmt")
	}
	fmt.Printf("\n%s", string(formatted))

	return nil
}
