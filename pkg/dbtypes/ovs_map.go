package dbtypes

import (
	"fmt"
	"reflect"

	jsoniter "github.com/json-iterator/go"
)

type OVSMapString map[string]string
type OVSMapInt map[string]int
type OVSMapUUID map[string]UUID

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
