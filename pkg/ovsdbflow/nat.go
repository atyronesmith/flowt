package ovsdbflow

type Nat struct {
	LogicalPort       string  `json:"logical_port,omitempty"`
	ExternalPortRange string  `json:"external_port_range"`
	ExternalIds       OVSMap[string]  `json:"external_ids"`
	Options           OVSMap[string]  `json:"options"`
	ExemptedExtIps    OVSMap[string]  `json:"exempted_ext_ips"`
	Type              NatType `json:"type"`
	AllowedExtIps     OVSMap[string]  `json:"allowed_ext_ips,omitempty"`
	ExternalIP        string  `json:"external_ip"`
	ExternalMac       string  `json:"external_mac,omitempty"`
	LogicalIP         string  `json:"logical_ip"`
}
