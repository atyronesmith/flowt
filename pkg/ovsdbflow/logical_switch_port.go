package ovsdbflow

type LogicalSwitchPort struct {
	Up          bool   `json:"up"`
	Name        string `json:"name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	// HaChassisGroup struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			Type     string `json:"type"`
	// 			RefTable string `json:"refTable"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"ha_chassis_group"`
	Options OVSMap[string] `json:"options"`
	// ParentName struct {
	// 	Type struct {
	// 		Min int    `json:"min"`
	// 		Key string `json:"key"`
	// 	} `json:"type"`
	// } `json:"parent_name"`
	Enabled bool `json:"enabled"`
	// Type struct {
	// 	Type string `json:"type"`
	// } `json:"type"`
	// Dhcpv6Options struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			RefType  string `json:"refType"`
	// 			Type     string `json:"type"`
	// 			RefTable string `json:"refTable"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"dhcpv6_options"`
	// TagRequest struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			MinInteger int    `json:"minInteger"`
	// 			MaxInteger int    `json:"maxInteger"`
	// 			Type       string `json:"type"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"tag_request"`
	//TODO
	PortSecurity string `json:"port_security"`
	// TODO
	Addresses string `json:"addresses"`
	// DynamicAddresses struct {
	// 	Type struct {
	// 		Min int    `json:"min"`
	// 		Key string `json:"key"`
	// 	} `json:"type"`
	// } `json:"dynamic_addresses"`
	// Dhcpv4Options struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			RefType  string `json:"refType"`
	// 			Type     string `json:"type"`
	// 			RefTable string `json:"refTable"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"dhcpv4_options"`
	// Tag struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			MinInteger int    `json:"minInteger"`
	// 			MaxInteger int    `json:"maxInteger"`
	// 			Type       string `json:"type"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"columns"`
}
