// Type: OVN_Northbound
// Version: 5.34.1
// Table Definitions: 27
//	DHCP_Options
//	Logical_Switch_Port
//	HA_Chassis_Group
//	Address_Set
//	Forwarding_Group
//	QoS
//	Port_Group
//	Meter_Band
//	NB_Global
//	Meter
//	Load_Balancer_Group
//	Load_Balancer_Health_Check
//	Gateway_Chassis
//	Logical_Router_Static_Route
//	Load_Balancer
//	DNS
//	Logical_Router
//	ACL
//	Connection
//	SSL
//	HA_Chassis
//	Copp
//	NAT
//	Logical_Router_Port
//	BFD
//	Logical_Router_Policy
//	Logical_Switch

package nb

import (
	types "github.com/atyronesmith/flowt/pkg/ovsdbflow"
)

type ACL struct {
	Action      string               `json:"action"`
	Direction   string               `json:"direction"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Label       int                  `json:"label"`
	Log         bool                 `json:"log"`
	Match       string               `json:"match"`
	Meter       string               `json:"meter"`
	Name        string               `json:"name"`
	Priority    int                  `json:"priority"`
	Severity    string               `json:"severity"`
}
type AddressSet struct {
	Addresses   types.OVSSet[string] `json:"addresses"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Name        string               `json:"name"`
}
type BFD struct {
	DetectMult  int                  `json:"detect_mult"`
	DstIp       string               `json:"dst_ip"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	LogicalPort string               `json:"logical_port"`
	MinRx       int                  `json:"min_rx"`
	MinTx       int                  `json:"min_tx"`
	Options     types.OVSMap[string] `json:"options"`
	Status      string               `json:"status"`
}
type Connection struct {
	ExternalIds     types.OVSMap[string] `json:"external_ids"`
	InactivityProbe int                  `json:"inactivity_probe"`
	IsConnected     bool                 `json:"is_connected"`
	MaxBackoff      int                  `json:"max_backoff"`
	OtherConfig     types.OVSMap[string] `json:"other_config"`
	Status          types.OVSMap[string] `json:"status"`
	Target          string               `json:"target"`
}
type Copp struct {
	Meters types.OVSMap[string] `json:"meters"`
}
type DHCPOptions struct {
	Cidr        string               `json:"cidr"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Options     types.OVSMap[string] `json:"options"`
}
type DNS struct {
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Records     types.OVSMap[string] `json:"records"`
}
type ForwardingGroup struct {
	ChildPort   types.OVSSet[string] `json:"child_port"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Liveness    bool                 `json:"liveness"`
	Name        string               `json:"name"`
	Vip         string               `json:"vip"`
	Vmac        string               `json:"vmac"`
}
type GatewayChassis struct {
	ChassisName string               `json:"chassis_name"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Name        string               `json:"name"`
	Options     types.OVSMap[string] `json:"options"`
	Priority    int                  `json:"priority"`
}
type HAChassis struct {
	ChassisName string               `json:"chassis_name"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Priority    int                  `json:"priority"`
}
type HAChassisGroup struct {
	ExternalIds types.OVSMap[string]     `json:"external_ids"`
	HaChassis   types.OVSSet[types.UUID] `json:"ha_chassis"`
	Name        string                   `json:"name"`
}
type LoadBalancer struct {
	ExternalIds     types.OVSMap[string]     `json:"external_ids"`
	HealthCheck     types.OVSSet[types.UUID] `json:"health_check"`
	IpPortMappings  types.OVSMap[string]     `json:"ip_port_mappings"`
	Name            string                   `json:"name"`
	Options         types.OVSMap[string]     `json:"options"`
	Protocol        string                   `json:"protocol"`
	SelectionFields types.OVSSet[string]     `json:"selection_fields"`
	Vips            types.OVSMap[string]     `json:"vips"`
}
type LoadBalancerGroup struct {
	LoadBalancer types.OVSSet[types.UUID] `json:"load_balancer"`
	Name         string                   `json:"name"`
}
type LoadBalancerHealthCheck struct {
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Options     types.OVSMap[string] `json:"options"`
	Vip         string               `json:"vip"`
}
type LogicalRouter struct {
	Copp              types.OVSSet[types.UUID] `json:"copp"`
	Enabled           bool                     `json:"enabled"`
	ExternalIds       types.OVSMap[string]     `json:"external_ids"`
	LoadBalancer      types.OVSSet[types.UUID] `json:"load_balancer"`
	LoadBalancerGroup types.OVSSet[types.UUID] `json:"load_balancer_group"`
	Name              string                   `json:"name"`
	Nat               types.OVSSet[types.UUID] `json:"nat"`
	Options           types.OVSMap[string]     `json:"options"`
	Policies          types.OVSSet[types.UUID] `json:"policies"`
	Ports             types.OVSSet[types.UUID] `json:"ports"`
	StaticRoutes      types.OVSSet[types.UUID] `json:"static_routes"`
}
type LogicalRouterPolicy struct {
	Action      string               `json:"action"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Match       string               `json:"match"`
	Nexthop     string               `json:"nexthop"`
	Nexthops    types.OVSSet[string] `json:"nexthops"`
	Options     types.OVSMap[string] `json:"options"`
	Priority    int                  `json:"priority"`
}
type LogicalRouterPort struct {
	Enabled        bool                     `json:"enabled"`
	ExternalIds    types.OVSMap[string]     `json:"external_ids"`
	GatewayChassis types.OVSSet[types.UUID] `json:"gateway_chassis"`
	HaChassisGroup types.OVSSet[types.UUID] `json:"ha_chassis_group"`
	Ipv6Prefix     types.OVSSet[string]     `json:"ipv6_prefix"`
	Ipv6RaConfigs  types.OVSMap[string]     `json:"ipv6_ra_configs"`
	Mac            string                   `json:"mac"`
	Name           string                   `json:"name"`
	Networks       types.OVSSet[string]     `json:"networks"`
	Options        types.OVSMap[string]     `json:"options"`
	Peer           string                   `json:"peer"`
}
type LogicalRouterStaticRoute struct {
	Bfd         types.OVSSet[types.UUID] `json:"bfd"`
	ExternalIds types.OVSMap[string]     `json:"external_ids"`
	IpPrefix    string                   `json:"ip_prefix"`
	Nexthop     string                   `json:"nexthop"`
	Options     types.OVSMap[string]     `json:"options"`
	OutputPort  string                   `json:"output_port"`
	Policy      string                   `json:"policy"`
	RouteTable  string                   `json:"route_table"`
}
type LogicalSwitch struct {
	Acls              types.OVSSet[types.UUID] `json:"acls"`
	Copp              types.OVSSet[types.UUID] `json:"copp"`
	DnsRecords        types.OVSSet[types.UUID] `json:"dns_records"`
	ExternalIds       types.OVSMap[string]     `json:"external_ids"`
	ForwardingGroups  types.OVSSet[types.UUID] `json:"forwarding_groups"`
	LoadBalancer      types.OVSSet[types.UUID] `json:"load_balancer"`
	LoadBalancerGroup types.OVSSet[types.UUID] `json:"load_balancer_group"`
	Name              string                   `json:"name"`
	OtherConfig       types.OVSMap[string]     `json:"other_config"`
	Ports             types.OVSSet[types.UUID] `json:"ports"`
	QosRules          types.OVSSet[types.UUID] `json:"qos_rules"`
}
type LogicalSwitchPort struct {
	Addresses        types.OVSSet[string]     `json:"addresses"`
	Dhcpv4Options    types.OVSSet[types.UUID] `json:"dhcpv4_options"`
	Dhcpv6Options    types.OVSSet[types.UUID] `json:"dhcpv6_options"`
	DynamicAddresses string                   `json:"dynamic_addresses"`
	Enabled          bool                     `json:"enabled"`
	ExternalIds      types.OVSMap[string]     `json:"external_ids"`
	HaChassisGroup   types.OVSSet[types.UUID] `json:"ha_chassis_group"`
	Name             string                   `json:"name"`
	Options          types.OVSMap[string]     `json:"options"`
	ParentName       string                   `json:"parent_name"`
	PortSecurity     types.OVSSet[string]     `json:"port_security"`
	Tag              int                      `json:"tag"`
	TagRequest       int                      `json:"tag_request"`
	Type             string                   `json:"type"`
	Up               bool                     `json:"up"`
}
type Meter struct {
	Bands       types.OVSSet[types.UUID] `json:"bands"`
	ExternalIds types.OVSMap[string]     `json:"external_ids"`
	Fair        bool                     `json:"fair"`
	Name        string                   `json:"name"`
	Unit        string                   `json:"unit"`
}
type MeterBand struct {
	Action      string               `json:"action"`
	BurstSize   int                  `json:"burst_size"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Rate        int                  `json:"rate"`
}
type NAT struct {
	AllowedExtIps     types.OVSSet[types.UUID] `json:"allowed_ext_ips"`
	ExemptedExtIps    types.OVSSet[types.UUID] `json:"exempted_ext_ips"`
	ExternalIds       types.OVSMap[string]     `json:"external_ids"`
	ExternalIp        string                   `json:"external_ip"`
	ExternalMac       string                   `json:"external_mac"`
	ExternalPortRange string                   `json:"external_port_range"`
	LogicalIp         string                   `json:"logical_ip"`
	LogicalPort       string                   `json:"logical_port"`
	Options           types.OVSMap[string]     `json:"options"`
	Type              string                   `json:"type"`
}
type NBGlobal struct {
	Connections    types.OVSSet[types.UUID] `json:"connections"`
	ExternalIds    types.OVSMap[string]     `json:"external_ids"`
	HvCfg          int                      `json:"hv_cfg"`
	HvCfgTimestamp int                      `json:"hv_cfg_timestamp"`
	Ipsec          bool                     `json:"ipsec"`
	Name           string                   `json:"name"`
	NbCfg          int                      `json:"nb_cfg"`
	NbCfgTimestamp int                      `json:"nb_cfg_timestamp"`
	Options        types.OVSMap[string]     `json:"options"`
	SbCfg          int                      `json:"sb_cfg"`
	SbCfgTimestamp int                      `json:"sb_cfg_timestamp"`
	Ssl            types.OVSSet[types.UUID] `json:"ssl"`
}
type PortGroup struct {
	Acls        types.OVSSet[types.UUID] `json:"acls"`
	ExternalIds types.OVSMap[string]     `json:"external_ids"`
	Name        string                   `json:"name"`
	Ports       types.OVSSet[types.UUID] `json:"ports"`
}
type QoS struct {
	Action      types.OVSMap[int]    `json:"action"`
	Bandwidth   types.OVSMap[int]    `json:"bandwidth"`
	Direction   string               `json:"direction"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Match       string               `json:"match"`
	Priority    int                  `json:"priority"`
}
type SSL struct {
	BootstrapCaCert bool                 `json:"bootstrap_ca_cert"`
	CaCert          string               `json:"ca_cert"`
	Certificate     string               `json:"certificate"`
	ExternalIds     types.OVSMap[string] `json:"external_ids"`
	PrivateKey      string               `json:"private_key"`
	SslCiphers      string               `json:"ssl_ciphers"`
	SslProtocols    string               `json:"ssl_protocols"`
}

type OVNNorthbound struct {
	ACL                      map[string]ACL                      `json:"ACL"`
	AddressSet               map[string]AddressSet               `json:"Address_Set"`
	BFD                      map[string]BFD                      `json:"BFD"`
	Connection               map[string]Connection               `json:"Connection"`
	Copp                     map[string]Copp                     `json:"Copp"`
	DHCPOptions              map[string]DHCPOptions              `json:"DHCP_Options"`
	DNS                      map[string]DNS                      `json:"DNS"`
	ForwardingGroup          map[string]ForwardingGroup          `json:"Forwarding_Group"`
	GatewayChassis           map[string]GatewayChassis           `json:"Gateway_Chassis"`
	HAChassis                map[string]HAChassis                `json:"HA_Chassis"`
	HAChassisGroup           map[string]HAChassisGroup           `json:"HA_Chassis_Group"`
	LoadBalancer             map[string]LoadBalancer             `json:"Load_Balancer"`
	LoadBalancerGroup        map[string]LoadBalancerGroup        `json:"Load_Balancer_Group"`
	LoadBalancerHealthCheck  map[string]LoadBalancerHealthCheck  `json:"Load_Balancer_Health_Check"`
	LogicalRouter            map[string]LogicalRouter            `json:"Logical_Router"`
	LogicalRouterPolicy      map[string]LogicalRouterPolicy      `json:"Logical_Router_Policy"`
	LogicalRouterPort        map[string]LogicalRouterPort        `json:"Logical_Router_Port"`
	LogicalRouterStaticRoute map[string]LogicalRouterStaticRoute `json:"Logical_Router_Static_Route"`
	LogicalSwitch            map[string]LogicalSwitch            `json:"Logical_Switch"`
	LogicalSwitchPort        map[string]LogicalSwitchPort        `json:"Logical_Switch_Port"`
	Meter                    map[string]Meter                    `json:"Meter"`
	MeterBand                map[string]MeterBand                `json:"Meter_Band"`
	NAT                      map[string]NAT                      `json:"NAT"`
	NBGlobal                 map[string]NBGlobal                 `json:"NB_Global"`
	PortGroup                map[string]PortGroup                `json:"Port_Group"`
	QoS                      map[string]QoS                      `json:"QoS"`
	SSL                      map[string]SSL                      `json:"SSL"`
}
