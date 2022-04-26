package ovsdbflow

type LogicalRouter struct {
	IsRoot            bool           `json:"isRoot"`
	Name              string         `json:"name"`
	Policies          OVSSet[string] `json:"policies"`
	ExternalIds       OVSMap[string]         `json:"external_ids"`
	Options           OVSMap[string]         `json:"options"`
	Enabled           bool           `json:"enabled,omitempty"`
	Nat               OVSSet[string] `json:"nat"`
	LoadBalancer      OVSSet[string] `json:"load_balancer"`
	LoadBalancerGroup OVSSet[string] `json:"load_balancer_group"`
	StaticRoutes      OVSSet[string] `json:"static_routes"`
	Ports             OVSSet[string] `json:"ports"`
	Copp              OVSSet[string] `json:"copp"`
}
