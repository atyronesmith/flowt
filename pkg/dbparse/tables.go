package dbparse

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"sort"

	types "github.com/atyronesmith/flowt/pkg/dbtypes"
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
		rType = "UUID"
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
		return fmt.Sprintf("OVSMap[%s]", rType), nil
	case Set:
		return fmt.Sprintf("OVSSet[%s]", rType), nil
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

func ParseSchema(tbl types.OVSdbSchema,pkg string) (*types.DbDef,error) {
	tblMap := tbl.Tables

	tableKeys := make([]string, 0, len(tblMap))
	for k := range tblMap {
		tableKeys = append(tableKeys, k)
	}

	sort.Strings(tableKeys)

	var db types.DbDef

	db.Name = utils.SnakeToCamel( tbl.Type.String() )

	for _, tblName := range tableKeys {
		var tblDef types.TableDef

		tblDefIntrf := tblMap[tblName]
		
		structName := utils.SnakeToCamel(tblName)

		tblDef.Name = structName
		tblDef.JsonName = tblName

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
						tblDef.Columns = append(tblDef.Columns, types.TableCol{ Name: utils.SnakeToCamel(colName), Type: pt, JsonName: colName})
					}
				}
			}
		}
		db.Tables = append(db.Tables, tblDef)
	}

	return &db, nil
}
