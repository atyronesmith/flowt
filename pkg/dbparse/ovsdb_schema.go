package dbparse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"reflect"
	"runtime"
	"sort"

	"github.com/atyronesmith/flowt/pkg/dbtypes"
	"github.com/atyronesmith/flowt/pkg/utils"
)

type TableCol struct {
	Name      string
	Type      string
	JsonName  string
	Optional  bool
	Ephemeral bool
	Index     bool
	Comment   string
	RefTable  string
	Range     []int
}

type TableDef struct {
	Name     string
	JsonName string
	Columns  []TableCol
	Indices  map[string]bool
	ToolTip  string
}

type DbDef struct {
	Schema *OVSdbSchema
	Name   string
	TableDefs []TableDef
}

type OVNDbType interface {
	IsValid() bool
}

type OVSdbSchema struct {
	ChkSum  string                 `json:"cksum"`
	Type    OVSDBType              `json:"name"`
	Version string                 `json:"version"`
	Tables  map[string]interface{} `json:"tables"`
}

func (schema *OVSdbSchema) ReadOvsDbSchemaFile(filename string) error {
	var in io.Reader

	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %s, %v", filename, err)
	}
	defer f.Close()

	in = f

	stats, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("error occured on file: %s, %v", filename, err)
	}

	return schema.ReadOvsDbSchema(in,int(stats.Size()))
}

func (schema *OVSdbSchema) ReadOvsDbSchemaBuf(buf bytes.Buffer) error {

	rdr := bytes.NewReader(buf.Bytes())

	return schema.ReadOvsDbSchema(rdr,buf.Len())
}

func (schema *OVSdbSchema) ReadOvsDbSchema(in io.Reader, maxTokenSize int) error {
	scanner := bufio.NewScanner(in)

	scanner.Buffer(make([]byte, 0), maxTokenSize)
	scanner.Split(bufio.ScanLines)

	// db file must be unformatted
	for scanner.Scan() {
		line := scanner.Text()

		if jsonObj.MatchString(line) {
			if err := json.Unmarshal([]byte(line), schema); err != nil {
				return fmt.Errorf("unable to unmarshall json string: %s, %v", line, err)
			}
			if len(schema.Version) > 0 && len(schema.ChkSum) > 0 && len(schema.Type.String()) > 0 {
				return nil
			}
			if schema.Type == Unknown {
				return fmt.Errorf("unknown ovsdb type %s", schema.Type)
			}
		}
	}

	return fmt.Errorf("error, could not find ovsdb header")
}

func (schema *OVSdbSchema) NewDb() (OVNDbType, error) {
	switch schema.Type {
	case NB:
		return &dbtypes.OVNNorthbound{}, nil
	case SB:
		return &dbtypes.OVNSouthbound{}, nil
	default:
		return nil, fmt.Errorf("unknown db type: %v", schema.Type)
	}
}

type CType int

const (
	Atomic CType = iota
	Map
	Set
)

func mapKey(key interface{}, t reflect.Kind, cType CType, column *TableCol) error {
	var aType string
	var keyMap map[string]interface{}

	switch t {
	case reflect.Map:
		keyMap = key.(map[string]interface{})
		aType = keyMap["type"].(string)
		if min, ok := keyMap["minInteger"].(int); ok {
			if max, ok := keyMap["maxInteger"].(int); ok {
				column.Range = make([]int, 2)
				column.Range[0] = min
				column.Range[1] = max
			}
		}
	case reflect.String:
		aType = key.(string)
	}

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
		if rTable, ok := keyMap["refTable"]; ok {
			column.RefTable = rTable.(string)
		}
	default:
		return fmt.Errorf("unknown type: %s", aType)
	}

	switch cType {
	case Atomic:
		column.Type = "*" + rType
		return nil
	case Map:
		column.Type = fmt.Sprintf("OVSMap[%s]", rType)
		return nil
	case Set:
		column.Type = fmt.Sprintf("OVSSet[%s]", rType)
		return nil
	}
	return fmt.Errorf("unknown type: %s", aType)
}

func parseType(t interface{}, column *TableCol) error {

	k := reflect.ValueOf(t).Kind()

	if k == reflect.String {
		// An atomic type, "integer", "real", "boolean", "string", or "uuid"
		return mapKey(t, k, Atomic, column)
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
				return mapKey(value, reflect.String, Map, column)
			} else if reflect.ValueOf(value).Kind() == reflect.Map {
				return mapKey(value, reflect.Map, Map, column)
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
				column.Optional = true
				return mapKey(key, reflect.ValueOf(key).Kind(), Atomic, column)
			} else if minInt == 1 && maxInt == 1 {
				return mapKey(key, reflect.ValueOf(key).Kind(), Atomic, column)
			} else if maxInt == math.MaxInt {
				return mapKey(key, reflect.ValueOf(key).Kind(), Set, column)
			} else {
				return fmt.Errorf("could not parse type, min: %d, max: %x", minInt, maxInt)
			}
		}

	}

	return fmt.Errorf("unable to parse type: %v", t)
}

func (db *DbDef) ParseSchema(tbl *OVSdbSchema) (error) {
	db.Schema = tbl

	tblMap := tbl.Tables // Tables contains the unmarshalled schema

	tableKeys := make([]string, 0, len(tblMap))
	for k := range tblMap {
		tableKeys = append(tableKeys, k)
	}

	sort.Strings(tableKeys)

	db.Name = utils.SnakeToCamel(tbl.Type.String())

	for _, tblName := range tableKeys {
		var tblDef TableDef

		tblDefIntrf := tblMap[tblName]

		structName := utils.SnakeToCamel(tblName)

		tblDef.Name = structName
		tblDef.JsonName = tblName
		tblDef.Indices = make(map[string]bool)
		tableDef := tblDefIntrf.(map[string]interface{})
		if colIntrf, ok := tableDef["indexes"]; ok {
			ind := colIntrf.([]interface{})[0].([]interface{})
			for _, v := range ind {
				tblDef.Indices[v.(string)] = true
			}
		}

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
				if typeIntrf, ok := colDef["type"]; ok {
					newCol := TableCol{}
					if err := parseType(typeIntrf, &newCol); err == nil {
						newCol.Name = utils.SnakeToCamel(colName)
						newCol.JsonName = colName
						ephemeral, hasEphemeral := colDef["ephemeral"]
						if hasEphemeral {
							newCol.Ephemeral = ephemeral.(bool)
						}
						if v, ok := tblDef.Indices[colName]; ok {
							newCol.Index = v
						}
						tblDef.Columns = append(tblDef.Columns, newCol)
					}
				}
			}
		}
		db.TableDefs = append(db.TableDefs, tblDef)
	}

	return nil
}

func (tbl *DbDef) AugmentSchema() error {
	var toolTipFile string

	switch tbl.Schema.Type {
	case NB:
		toolTipFile = "cmd/db/data/nb_table_tooltip.json"
	case SB:
		toolTipFile = "cmd/db/data/sb_table_tooltip.json"
	}

	dat, err := os.ReadFile(toolTipFile)
	if err != nil {
		return fmt.Errorf("error reading file: %s, %v", toolTipFile, err)
	}

	toolTipObj := make(map[string]string)
	if err := json.Unmarshal(dat, &toolTipObj); err != nil {
		return fmt.Errorf("unable to unmarshall json string: %v", err)
	}

	for toolKey,toolValue := range toolTipObj {
		for idx := range tbl.TableDefs {
			if toolKey == tbl.TableDefs[idx].JsonName {
				tbl.TableDefs[idx].ToolTip = toolValue
			}
		}
	}

	return nil
}
