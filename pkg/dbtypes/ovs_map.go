package dbtypes

import (
	"fmt"
	"reflect"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

type OVSMapString map[string]string
type OVSMapInt map[string]int
type OVSMapUUID map[string]UUID

type LogFlow LogicalFlowSB

func (b *LogFlow) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	 for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
		switch field {
		case "match":
			match := iter.ReadString()
			(*LogFlow)(ptr).Match = &match
		case "pipeline":
			pipeline := iter.ReadString()
			(*LogFlow)(ptr).Pipeline = &pipeline
		case "actions":
			actions := iter.ReadString()
			(*LogFlow)(ptr).Actions = &actions
		case "controller_meter":
			meter := iter.ReadString()
			(*LogFlow)(ptr).ControllerMeter = &meter
		case "priority":
			priority := iter.ReadInt()
			(*LogFlow)(ptr).Priority = &priority
		case "table_id":
			id := iter.ReadInt()
			(*LogFlow)(ptr).TableId = &id
		case "logical_datapath":
			// ["uuid","304230cf-886d-45eb-8548-39b6f9b926fc"]
			iter.ReadArray()
			iter.ReadString()
			iter.ReadArray()
			uuid := UUID(iter.ReadString())
			(*LogFlow)(ptr).LogicalDatapath = &uuid
			iter.ReadArray()
		case "logical_dp_group":
			// ["uuid","304230cf-886d-45eb-8548-39b6f9b926fc"]
			iter.ReadArray()  // "["
			iter.ReadString() // "uuid"
			iter.ReadArray()  // ","
			uuid := UUID(iter.ReadString()) // "304230cf-886d-45eb-8548-39b6f9b926fc"
			(*LogFlow)(ptr).LogicalDpGroup = &uuid
			iter.ReadArray()
		case "external_ids":
        //    "external_ids": [
        //         "map",
        //         [
        //             [
        //                 "source",
        //                 "northd.c:12675"
        //             ],
		// ...
        //             [
        //                 "stage-name",
        //                 "lr_out_chk_dnat_local"
        //             ]
        //         ]
        //     ]
			(*LogFlow)(ptr).ExternalIds = make(map[string]string)

			iter.ReadArray()  // "["
			iter.ReadString() // "map"
			iter.ReadArray()  // ","
			iter.ReadArray()  // "["
			for iter.ReadArray() {
				key := iter.ReadString() // key
				iter.ReadArray() // ","
				value := iter.ReadString() // value
				(*LogFlow)(ptr).ExternalIds[key] = value
				iter.ReadArray() //"]"
				iter.ReadArray() //","
			}
		default:
			fmt.Printf("Unknown field (%s) in LogicalFlowSB!\n",field)
			iter.Skip()
		}
	 }	
}

func (b *OVSMapString) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}
//	fmt.Printf("String -- %s\n",data)

	if err := jsoniter.Unmarshal(data, &ovsMap); err != nil {
		return fmt.Errorf("error unmarshaling OVSMap: %v",err)
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array for OVSMap: %s", data)
	}

	*b = make(map[string]string)

	for _, v := range ovsMap[1].([]interface{}) {
		(*b)[v.([]interface{})[0].(string)] = v.([]interface{})[1].(string)
	}

	return nil
}

func (b *OVSMapInt) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

//	fmt.Printf("Int -- %s\n",data)

	if err := jsoniter.Unmarshal(data, &ovsMap); err != nil {
		return fmt.Errorf("error unmarshaling OVSMap: %v",err)
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array for OVSMap: %s", data)
	}

	*b = make(map[string]int)

	for _, v := range ovsMap[1].([]interface{}) {
		(*b)[v.([]interface{})[0].(string)] = v.([]interface{})[1].(int)
	}

	return nil
}

func (b *OVSMapUUID) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

	fmt.Printf("UUID -- %s\n",data)

	if err := jsoniter.Unmarshal(data, &ovsMap); err != nil {
		return fmt.Errorf("error unmarshaling OVSMap: %v",err)
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array for OVSMap: %s", data)
	}

	*b = make(map[string]UUID)

	for _, v := range ovsMap[1].([]interface{}) {
		if reflect.ValueOf(v.([]interface{})[1]).Kind() == reflect.Slice {
			s := v.([]interface{})[1].([]interface{})
			(*b)[v.([]interface{})[0].(string)] = UUID(s[1].(string))
		} else {
			(*b)[v.([]interface{})[0].(string)] = UUID(v.([]interface{})[1].(string))
		}
	}

	fmt.Printf("UUID -- %s\n",*b)

	return nil
}

// func (b *OVSMap[T]) append(t []interface{}) error {
// 	var ret T
// 	switch p := any(&ret).(type) {
// 	case *UUID:
// 		if reflect.ValueOf(t[1]).Kind() == reflect.Slice {
// 			s := t[1].([]interface{})
// 			*p = UUID(s[1].(string))
// 		} else {
// 			*p = UUID(t[1].(string))
// 		}
// 		os.Exit(0)
// 	case *string:
// 		*p = t[1].(string)
// 	default:
// 		return fmt.Errorf("unknown type")
// 	}
// 	(*b)[t[0].(string)] = ret

// 	return nil
// }
