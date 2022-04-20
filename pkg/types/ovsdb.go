package types

type OVSdb struct {
	LogicalRouterStaticRoute bool                   `json:"Logical_Router_Static_Route"`
	NBGlobal                 string                 `json:"NB_Global"`
	LogicalSwitch            OVSMap                 `json:"Logical_Switch"`
	ACL                      bool                   `json:"ACL"`
	LogicalSwitchPortTable   LogicalSwitchPortTable `json:"Logical_Switch_Port"`
}
