package types

import (
	"encoding/json"
	"fmt"
)

type OVSMap map[string]string

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

func (b *OVSMap) UnmarshalJSON(data []byte) error {
	//	type Alias OVSMap

	//	fmt.Printf("Unmarshal: %s\n", data)

	var ovsMap []interface{}

	if err := json.Unmarshal(data, &ovsMap); err != nil {
		return err
	}

	if len(ovsMap) != 2 {
		return fmt.Errorf("invalid map array: %s", data)
	}
	
	*b = make(map[string]string)

	for _, v := range ovsMap[1].([]interface{}) {
		(*b)[v.([]interface{})[0].(string)] = v.([]interface{})[1].(string)
	}

	return nil
}

