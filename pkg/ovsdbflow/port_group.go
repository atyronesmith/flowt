package ovsdbflow

type PortGroup struct {
	IsRoot      bool           `json:"isRoot"`
	Name        string         `json:"name"`
	Ports       OVSSet[string] `json:"ports"`
	ExternalIds OVSMap[string]         `json:"external_ids"`
	Acls        OVSSet[string] `json:"acls"`
}
