package ovsdbflow

type LogicalSwitch struct {
	IsRoot            bool           `json:"isRoot"`
	Name              string         `json:"name"`
	QosRules          OVSSet[string] `json:"qos_rules"`
	ExternalIds       OVSMap[string] `json:"external_ids"`
	OtherConfig       OVSMap[string] `json:"other_config"`
	DNSRecords        OVSSet[string] `json:"dns_records"`
	Ports             OVSSet[string] `json:"ports"`
	LoadBalancer      OVSSet[string] `json:"load_balancer"`
	LoadBalancerGroup OVSSet[string] `json:"load_balancer_group"`
	ForwardingGroups  OVSSet[string] `json:"forwarding_groups"`
	Copp              OVSSet[string] `json:"copp"`
	Acls              OVSSet[string] `json:"acls"`
}
