package ovsdbflow

import (
	"encoding/json"
	"fmt"
)

type OVSMap[T string|int] map[string]T

//    "options": [
//        "map",
//        [
//            [
//                "mcast_flood_reports",
//                "true"
//            ],
//            [
//                "requested-chassis",
//                "compute1029utn10rt-0.redhat.local"
//            ]
//        ]
//    ],

func (b *OVSMap[T]) UnmarshalJSON(data []byte) error {
	var ovsMap []interface{}

	if err := json.Unmarshal(data, &ovsMap); err != nil {
					fmt.Printf("%s\n",string(data))

		return err
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array: %s", data)
	}
	
	*b = make(map[string]T)

	for _, v := range ovsMap[1].([]interface{}) {
		(*b)[v.([]interface{})[0].(string)] = v.([]interface{})[1].(T)
	}

	return nil
}

