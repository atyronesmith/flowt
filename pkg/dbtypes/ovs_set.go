package dbtypes

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var strRegex = regexp.MustCompile(`^\".*\"$`)

type OVSSet[T string | int | UUID] []T

func (b *OVSSet[T]) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

	if strRegex.Match(data) {
		var o interface{} = strings.ReplaceAll(string(data), "\"", "")

		(*b) = append((*b), o.(T))
	} else {
		if err := jsoniter.Unmarshal(data, &ovsMap); err != nil {
			return fmt.Errorf("error while marshaling for OVSSet: %v",err)
		}
		if reflect.ValueOf(ovsMap[1]).Kind() == reflect.String {
			if len(ovsMap) != 2 {
				return fmt.Errorf("cannot unmarshal %s into OVSSet", string(data))
			}
			b.append(ovsMap[1])
		} else {
			var setSlice []interface{} = ovsMap[1].([]interface{})
			for _, v := range setSlice {
				if reflect.ValueOf(v).Kind() == reflect.Slice {
					var set []interface{} = v.([]interface{})

					b.append(set[1])
				} else if reflect.ValueOf(v).Kind() != reflect.Slice {
					b.append(v)
				}
			}
		}
	}
	return nil
}

func (b *OVSSet[T]) append(t interface{}) error {
	var ret T
	switch p := any(&ret).(type) {
	case *UUID:
		*p = UUID(t.(string))
	case *string:
		*p = t.(string)
	default:
		return fmt.Errorf("unknown type")
	}
	(*b) = append((*b), ret)

	return nil
}
