package dbtypes

// Type: OVN_Northbound
// Version: 5.34.1
// Tables: 27

// ACL
type ACLNB struct {
	Action      string         `json:"action"`
	Direction   string         `json:"direction"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Label       int            `json:"label"`
	Log         bool           `json:"log"`
	Match       string         `json:"match"`
	Meter       string         `json:"meter"`
	Name        string         `json:"name"`
	Priority    int            `json:"priority"`
	Severity    string         `json:"severity"`
}

// Address_Set
type AddressSetNB struct {
	Addresses   OVSSet[string] `json:"addresses"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        string         `json:"name"`
}

// BFD
type BFDNB struct {
	DetectMult  int            `json:"detect_mult"`
	DstIp       string         `json:"dst_ip"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	LogicalPort string         `json:"logical_port"`
	MinRx       int            `json:"min_rx"`
	MinTx       int            `json:"min_tx"`
	Options     OVSMap[string] `json:"options"`
	Status      string         `json:"status"`
}

// Connection
type ConnectionNB struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	InactivityProbe int            `json:"inactivity_probe"`
	IsConnected     bool           `json:"is_connected"`
	MaxBackoff      int            `json:"max_backoff"`
	OtherConfig     OVSMap[string] `json:"other_config"`
	Status          OVSMap[string] `json:"status"`
	Target          string         `json:"target"`
}

// Copp
type CoppNB struct {
	Meters OVSMap[string] `json:"meters"`
}

// DHCP_Options
type DHCPOptionsNB struct {
	Cidr        string         `json:"cidr"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Options     OVSMap[string] `json:"options"`
}

// DNS
type DNSNB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Records     OVSMap[string] `json:"records"`
}

// Forwarding_Group
type ForwardingGroupNB struct {
	ChildPort   OVSSet[string] `json:"child_port"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Liveness    bool           `json:"liveness"`
	Name        string         `json:"name"`
	Vip         string         `json:"vip"`
	Vmac        string         `json:"vmac"`
}

// Gateway_Chassis
type GatewayChassisNB struct {
	ChassisName string         `json:"chassis_name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        string         `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Priority    int            `json:"priority"`
}

// HA_Chassis
type HAChassisNB struct {
	ChassisName string         `json:"chassis_name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Priority    int            `json:"priority"`
}

// HA_Chassis_Group
type HAChassisGroupNB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	HaChassis   OVSSet[UUID]   `json:"ha_chassis"`
	Name        string         `json:"name"`
}

// Load_Balancer
type LoadBalancerNB struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	HealthCheck     OVSSet[UUID]   `json:"health_check"`
	IpPortMappings  OVSMap[string] `json:"ip_port_mappings"`
	Name            string         `json:"name"`
	Options         OVSMap[string] `json:"options"`
	Protocol        string         `json:"protocol"`
	SelectionFields OVSSet[string] `json:"selection_fields"`
	Vips            OVSMap[string] `json:"vips"`
}

// Load_Balancer_Group
type LoadBalancerGroupNB struct {
	LoadBalancer OVSSet[UUID] `json:"load_balancer"`
	Name         string       `json:"name"`
}

// Load_Balancer_Health_Check
type LoadBalancerHealthCheckNB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Options     OVSMap[string] `json:"options"`
	Vip         string         `json:"vip"`
}

// Logical_Router
type LogicalRouterNB struct {
	Copp              OVSSet[UUID]   `json:"copp"`
	Enabled           bool           `json:"enabled"`
	ExternalIds       OVSMap[string] `json:"external_ids"`
	LoadBalancer      OVSSet[UUID]   `json:"load_balancer"`
	LoadBalancerGroup OVSSet[UUID]   `json:"load_balancer_group"`
	Name              string         `json:"name"`
	Nat               OVSSet[UUID]   `json:"nat"`
	Options           OVSMap[string] `json:"options"`
	Policies          OVSSet[UUID]   `json:"policies"`
	Ports             OVSSet[UUID]   `json:"ports"`
	StaticRoutes      OVSSet[UUID]   `json:"static_routes"`
}

// Logical_Router_Policy
type LogicalRouterPolicyNB struct {
	Action      string         `json:"action"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Match       string         `json:"match"`
	Nexthop     string         `json:"nexthop"`
	Nexthops    OVSSet[string] `json:"nexthops"`
	Options     OVSMap[string] `json:"options"`
	Priority    int            `json:"priority"`
}

// Logical_Router_Port
type LogicalRouterPortNB struct {
	Enabled        bool           `json:"enabled"`
	ExternalIds    OVSMap[string] `json:"external_ids"`
	GatewayChassis OVSSet[UUID]   `json:"gateway_chassis"`
	HaChassisGroup OVSSet[UUID]   `json:"ha_chassis_group"`
	Ipv6Prefix     OVSSet[string] `json:"ipv6_prefix"`
	Ipv6RaConfigs  OVSMap[string] `json:"ipv6_ra_configs"`
	Mac            string         `json:"mac"`
	Name           string         `json:"name"`
	Networks       OVSSet[string] `json:"networks"`
	Options        OVSMap[string] `json:"options"`
	Peer           string         `json:"peer"`
}

// Logical_Router_Static_Route
type LogicalRouterStaticRouteNB struct {
	Bfd         OVSSet[UUID]   `json:"bfd"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	IpPrefix    string         `json:"ip_prefix"`
	Nexthop     string         `json:"nexthop"`
	Options     OVSMap[string] `json:"options"`
	OutputPort  string         `json:"output_port"`
	Policy      string         `json:"policy"`
	RouteTable  string         `json:"route_table"`
}

// Logical_Switch
type LogicalSwitchNB struct {
	Acls              OVSSet[UUID]   `json:"acls"`
	Copp              OVSSet[UUID]   `json:"copp"`
	DnsRecords        OVSSet[UUID]   `json:"dns_records"`
	ExternalIds       OVSMap[string] `json:"external_ids"`
	ForwardingGroups  OVSSet[UUID]   `json:"forwarding_groups"`
	LoadBalancer      OVSSet[UUID]   `json:"load_balancer"`
	LoadBalancerGroup OVSSet[UUID]   `json:"load_balancer_group"`
	Name              string         `json:"name"`
	OtherConfig       OVSMap[string] `json:"other_config"`
	Ports             OVSSet[UUID]   `json:"ports"`
	QosRules          OVSSet[UUID]   `json:"qos_rules"`
}

// Logical_Switch_Port
type LogicalSwitchPortNB struct {
	Addresses        OVSSet[string] `json:"addresses"`
	Dhcpv4Options    OVSSet[UUID]   `json:"dhcpv4_options"`
	Dhcpv6Options    OVSSet[UUID]   `json:"dhcpv6_options"`
	DynamicAddresses string         `json:"dynamic_addresses"`
	Enabled          bool           `json:"enabled"`
	ExternalIds      OVSMap[string] `json:"external_ids"`
	HaChassisGroup   OVSSet[UUID]   `json:"ha_chassis_group"`
	Name             string         `json:"name"`
	Options          OVSMap[string] `json:"options"`
	ParentName       string         `json:"parent_name"`
	PortSecurity     OVSSet[string] `json:"port_security"`
	Tag              int            `json:"tag"`
	TagRequest       int            `json:"tag_request"`
	Type             string         `json:"type"`
	Up               bool           `json:"up"`
}

// Meter
type MeterNB struct {
	Bands       OVSSet[UUID]   `json:"bands"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Fair        bool           `json:"fair"`
	Name        string         `json:"name"`
	Unit        string         `json:"unit"`
}

// Meter_Band
type MeterBandNB struct {
	Action      string         `json:"action"`
	BurstSize   int            `json:"burst_size"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Rate        int            `json:"rate"`
}

// NAT
type NATNB struct {
	AllowedExtIps     OVSSet[UUID]   `json:"allowed_ext_ips"`
	ExemptedExtIps    OVSSet[UUID]   `json:"exempted_ext_ips"`
	ExternalIds       OVSMap[string] `json:"external_ids"`
	ExternalIp        string         `json:"external_ip"`
	ExternalMac       string         `json:"external_mac"`
	ExternalPortRange string         `json:"external_port_range"`
	LogicalIp         string         `json:"logical_ip"`
	LogicalPort       string         `json:"logical_port"`
	Options           OVSMap[string] `json:"options"`
	Type              string         `json:"type"`
}

// NB_Global
type NBGlobalNB struct {
	Connections    OVSSet[UUID]   `json:"connections"`
	ExternalIds    OVSMap[string] `json:"external_ids"`
	HvCfg          int            `json:"hv_cfg"`
	HvCfgTimestamp int            `json:"hv_cfg_timestamp"`
	Ipsec          bool           `json:"ipsec"`
	Name           string         `json:"name"`
	NbCfg          int            `json:"nb_cfg"`
	NbCfgTimestamp int            `json:"nb_cfg_timestamp"`
	Options        OVSMap[string] `json:"options"`
	SbCfg          int            `json:"sb_cfg"`
	SbCfgTimestamp int            `json:"sb_cfg_timestamp"`
	Ssl            OVSSet[UUID]   `json:"ssl"`
}

// Port_Group
type PortGroupNB struct {
	Acls        OVSSet[UUID]   `json:"acls"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        string         `json:"name"`
	Ports       OVSSet[UUID]   `json:"ports"`
}

// QoS
type QoSNB struct {
	Action      OVSMap[int]    `json:"action"`
	Bandwidth   OVSMap[int]    `json:"bandwidth"`
	Direction   string         `json:"direction"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Match       string         `json:"match"`
	Priority    int            `json:"priority"`
}

// SSL
type SSLNB struct {
	BootstrapCaCert bool           `json:"bootstrap_ca_cert"`
	CaCert          string         `json:"ca_cert"`
	Certificate     string         `json:"certificate"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	PrivateKey      string         `json:"private_key"`
	SslCiphers      string         `json:"ssl_ciphers"`
	SslProtocols    string         `json:"ssl_protocols"`
}

type OVNNorthbound struct {
	Date                     Time                                  `json:"_date"`
	Comment                  string                                `json:"_comment"`
	IsDiff                   bool                                  `json:"_is_diff"`
	ACL                      map[string]ACLNB                      `json:"ACL"`
	AddressSet               map[string]AddressSetNB               `json:"Address_Set"`
	BFD                      map[string]BFDNB                      `json:"BFD"`
	Connection               map[string]ConnectionNB               `json:"Connection"`
	Copp                     map[string]CoppNB                     `json:"Copp"`
	DHCPOptions              map[string]DHCPOptionsNB              `json:"DHCP_Options"`
	DNS                      map[string]DNSNB                      `json:"DNS"`
	ForwardingGroup          map[string]ForwardingGroupNB          `json:"Forwarding_Group"`
	GatewayChassis           map[string]GatewayChassisNB           `json:"Gateway_Chassis"`
	HAChassis                map[string]HAChassisNB                `json:"HA_Chassis"`
	HAChassisGroup           map[string]HAChassisGroupNB           `json:"HA_Chassis_Group"`
	LoadBalancer             map[string]LoadBalancerNB             `json:"Load_Balancer"`
	LoadBalancerGroup        map[string]LoadBalancerGroupNB        `json:"Load_Balancer_Group"`
	LoadBalancerHealthCheck  map[string]LoadBalancerHealthCheckNB  `json:"Load_Balancer_Health_Check"`
	LogicalRouter            map[string]LogicalRouterNB            `json:"Logical_Router"`
	LogicalRouterPolicy      map[string]LogicalRouterPolicyNB      `json:"Logical_Router_Policy"`
	LogicalRouterPort        map[string]LogicalRouterPortNB        `json:"Logical_Router_Port"`
	LogicalRouterStaticRoute map[string]LogicalRouterStaticRouteNB `json:"Logical_Router_Static_Route"`
	LogicalSwitch            map[string]LogicalSwitchNB            `json:"Logical_Switch"`
	LogicalSwitchPort        map[string]LogicalSwitchPortNB        `json:"Logical_Switch_Port"`
	Meter                    map[string]MeterNB                    `json:"Meter"`
	MeterBand                map[string]MeterBandNB                `json:"Meter_Band"`
	NAT                      map[string]NATNB                      `json:"NAT"`
	NBGlobal                 map[string]NBGlobalNB                 `json:"NB_Global"`
	PortGroup                map[string]PortGroupNB                `json:"Port_Group"`
	QoS                      map[string]QoSNB                      `json:"QoS"`
	SSL                      map[string]SSLNB                      `json:"SSL"`
}

func (nb OVNNorthbound) IsValid() bool {
	return len(nb.LogicalSwitchPort) > 0
}
