package ovsdbflow

// Each row in this table represents one ACL rule for a logical switch  or
// a port group that points to it through its acls column. The action col?
// umn for the highest-priority matching row in this  table  determines  a
// packet?s  treatment. If no row matches, packets are allowed by default.
// (Default-deny treatment is possible: add a rule with priority 0,  1  as
// match, and deny as action.)

type ACL struct {
	Name string `json:"name"`
	// integer, in range 0 to 32,767
	Priority    uint   `json:"priority"`
	Log         bool   `json:"log"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	// string, either from-lport or to-lport
	Direction AclDirection `json:"direction"`
	// Meter     struct {
	// 	Type struct {
	// 		Min int    `json:"min"`
	// 		Key string `json:"key"`
	// 	} `json:"type"`
	// } `json:"meter"`
	//  integer, in range 0 to 4,294,967,295
	Label  uint      `json:"label"`
	Action AclAction `json:"action"`
	// The  packets  that  the ACL should match, in the same expression
	// language used for the match column in the OVN  Southbound  data?
	// base?s  Logical_Flow  table.
	Match string `json:"match"`
	// Severity struct {
	// 	Type struct {
	// 		Min int `json:"min"`
	// 		Key struct {
	// 			Type string        `json:"type"`
	// 			Enum []interface{} `json:"enum"`
	// 		} `json:"key"`
	// 	} `json:"type"`
	// } `json:"severity"`
}
