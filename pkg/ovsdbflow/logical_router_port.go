package ovsdbflow

type LogicalRouterPort struct {
	Name           string         `json:"name"`
	Ipv6Prefix     OVSSet[string] `json:"ipv6_prefix"`
	Mac            string         `json:"mac"`
	ExternalIds    OVSMap[string]         `json:"external_ids"`
	GatewayChassis OVSSet[string] `json:"gateway_chassis"`
	HaChassisGroup OVSSet[string] `json:"ha_chassis_group"`
	Options        OVSMap[string]         `json:"options"`
	Enabled        bool           `json:"enabled"`
	Networks       string         `json:"networks"`
	Peer           string         `json:"peer"`
	Ipv6RaConfigs  OVSMap[string]        `json:"ipv6_ra_configs"`
}
