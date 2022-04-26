package ovsdbflow

type GatewayChassis struct {
	Name        string `json:"name"`
	Priority    uint   `json:"priority"`
	ChassisName string `json:"chassis_name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Options     OVSMap[string] `json:"options"`
}
