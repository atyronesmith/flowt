package ovsdbflow

// Table representing a group of chassis which can provide high availabil?
// ity services. Each chassis in the group is  represented  by  the  table
// HA_Chassis.  The HA chassis with highest priority will be the master of
// this group. If the master chassis failover is detected, the HA  chassis
// with  the next higher priority takes over the responsibility of provid?
// ing the HA. If a distributed gateway router port references  a  row  in
// this table, then the master HA chassis in this group provides the gate?
// way functionality.

type HAChassisGroup struct {
	IsRoot      bool       `json:"isRoot"`
	Indexes     [][]string `json:"indexes"`
	Name        string     `json:"name"`
	ExternalIds OVSMap[string]     `json:"external_ids"`
	HaChassis   OVSMap[string]     `json:"ha_chassis"`
}
