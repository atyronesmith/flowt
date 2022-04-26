package ovsdbflow

type LogicalRouterStaticRoute struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Options     OVSMap[string] `json:"options"`
	// Bfd struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			RefType  string `json:"refType"`
	// 			Type     string `json:"type"`
	// 			RefTable string `json:"refTable"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"bfd"`
	IPPrefix   string `json:"ip_prefix"`
	RouteTable string `json:"route_table"`
	// Policy struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			Type string        `json:"type"`
	// 			Enum []interface{} `json:"enum"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"policy"`
	Nexthop    string `json:"nexthop"`
	OutputPort string `json:"output_port"`
}
