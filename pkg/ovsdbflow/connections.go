package ovsdbflow

type Connection struct {
	IsConnected     bool           `json:"is_connected"`
	Target          string         `json:"target"`
	OtherConfig     OVSMap[string] `json:"other_config"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	Status          OVSMap[string] `json:"status"`
	MaxBackoff      uint           `json:"max_backoff"`
	InactivityProbe uint           `json:"inactivity_probe"`
}
