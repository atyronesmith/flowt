package dbtypes

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type OVSMap[T string | int | UUID] map[string]T

func (b *OVSMap[T]) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

	if err := json.Unmarshal(data, &ovsMap); err != nil {
		return fmt.Errorf("error unmarshaling OVSMap: %v",err)
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array for OVSMap: %s", data)
	}

	*b = make(map[string]T,4)

	for _, v := range ovsMap[1].([]interface{}) {
		b.append(v.([]interface{}))
	}

	return nil
}

func (b *OVSMap[T]) append(t []interface{}) error {
	var ret T
	switch p := any(&ret).(type) {
	case *UUID:
		if reflect.ValueOf(t[1]).Kind() == reflect.Slice {
			s := t[1].([]interface{})
			*p = UUID(s[1].(string))
		} else {
			*p = UUID(t[1].(string))
		}
	case *string:
		*p = t[1].(string)
	default:
		return fmt.Errorf("unknown type")
	}
	(*b)[t[0].(string)] = ret

	return nil
}
