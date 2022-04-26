package ovsdbflow

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var strRegex = regexp.MustCompile(`^\".*\"$`)
//var arrRegex = regexp.MustCompile(`^\[.*\]$`)

type OVSSet[T ~string | int] []T

// "nat": {
//     "type": {
//         "max": "unlimited",
//         "min": 0,
//         "key": {
//             "type": "uuid",
//             "refTable": "NAT"
//         }
//     }
// },
//"nat": [
//    "set",
//    [
//        [
//            "uuid",
//            "006f6727-673f-439a-80d5-67897a2d79f1"
//        ],
//        [
//            "uuid",
//            "030b5d36-1451-4ca2-8305-f4e75962d492"
//        ],
//        [
//            "uuid",
//            "03ad39b6-fbbb-4728-adbd-d7bcd274d494"
//        ],
//     ]
// ]

func (b *OVSSet[T]) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

	if strRegex.Match(data) {
		var o interface{} = strings.ReplaceAll(string(data),"\"","")

		(*b) = append((*b), o.(T))
	} else {
		if err := json.Unmarshal(data, &ovsMap); err != nil {
			fmt.Printf("%s\n",string(data))
			return err
		}
		if reflect.ValueOf(ovsMap[1]).Kind() == reflect.String {
			if len(ovsMap) != 2 {
				return fmt.Errorf("cannot unmarshal %s into OVSSet", string(data))
			}
			(*b) = append((*b), ovsMap[1].(T))
		} else {
			var setSlice []interface{} = ovsMap[1].([]interface{})

			for _, v := range setSlice {
				if reflect.ValueOf(v).Kind() != reflect.Slice {
					return fmt.Errorf("cannot unmarshal OVSSet: %v ", v)
				}
				var set []interface{} = v.([]interface{})
				(*b) = append((*b), set[1].(T))
			}
		}
	}
	return nil
}
