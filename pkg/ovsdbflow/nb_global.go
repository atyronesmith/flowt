package ovsdbflow

type NBGlobal struct {
	MaxRows        int    `json:"maxRows"`
	IsRoot         bool   `json:"isRoot"`
	SbCfgTimestamp uint   `json:"sb_cfg_timestamp"`
	Name           string `json:"name"`
	HvCfg          uint   `json:"hv_cfg"`
	NbCfg          uint   `json:"nb_cfg"`
	ExternalIds    OVSMap[string] `json:"external_ids"`
	Options        OVSMap[string] `json:"options"`
	SbCfg          uint   `json:"sb_cfg"`
	// Ssl struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			Type     string `json:"type"`
	// 			RefTable string `json:"refTable"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"ssl"`
	Ipsec          string   `json:"ipsec"`
	HvCfgTimestamp uint     `json:"hv_cfg_timestamp"`
	Connections    []string `json:"connections"`
	NbCfgTimestamp uint     `json:"nb_cfg_timestamp"`
}

