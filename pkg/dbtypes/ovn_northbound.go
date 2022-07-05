package dbtypes

// Type: OVN_Northbound
// Version: 5.35.1
// Tables: 27

// ACL
type ACLNB struct {
	Action      *string        `json:"action"`
	Direction   *string        `json:"direction"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Label       *int           `json:"label"`
	Log         *bool          `json:"log"`
	Match       *string        `json:"match"`
	Meter       *string        `json:"meter,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Options     OVSMap[string] `json:"options"`
	Priority    *int           `json:"priority"`
	Severity    *string        `json:"severity,omitempty"`
}

// Address_Set
type AddressSetNB struct {
	Addresses   OVSSet[string] `json:"addresses"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        *string        `json:"name"`
}

// BFD
type BFDNB struct {
	DetectMult  *int           `json:"detect_mult,omitempty"`
	DstIp       *string        `json:"dst_ip"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	LogicalPort *string        `json:"logical_port"`
	MinRx       *int           `json:"min_rx,omitempty"`
	MinTx       *int           `json:"min_tx,omitempty"`
	Options     OVSMap[string] `json:"options"`
	Status      *string        `json:"status,omitempty"`
}

// Connection
type ConnectionNB struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	InactivityProbe *int           `json:"inactivity_probe,omitempty"`
	IsConnected     *bool          `json:"is_connected"`
	MaxBackoff      *int           `json:"max_backoff,omitempty"`
	OtherConfig     OVSMap[string] `json:"other_config"`
	Status          OVSMap[string] `json:"status"`
	Target          *string        `json:"target"`
}

// Copp
type CoppNB struct {
	Meters OVSMap[string] `json:"meters"`
}

// DHCP_Options
type DHCPOptionsNB struct {
	Cidr        *string        `json:"cidr"`
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
	Liveness    *bool          `json:"liveness"`
	Name        *string        `json:"name"`
	Vip         *string        `json:"vip"`
	Vmac        *string        `json:"vmac"`
}

// Gateway_Chassis
type GatewayChassisNB struct {
	ChassisName *string        `json:"chassis_name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        *string        `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Priority    *int           `json:"priority"`
}

// HA_Chassis
type HAChassisNB struct {
	ChassisName *string        `json:"chassis_name"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Priority    *int           `json:"priority"`
}

// HA_Chassis_Group
type HAChassisGroupNB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	HaChassis   OVSSet[UUID]   `json:"ha_chassis"` //  HA_Chassis
	Name        *string        `json:"name"`
}

// Load_Balancer
type LoadBalancerNB struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	HealthCheck     OVSSet[UUID]   `json:"health_check"` //  Load_Balancer_Health_Check
	IpPortMappings  OVSMap[string] `json:"ip_port_mappings"`
	Name            *string        `json:"name"`
	Options         OVSMap[string] `json:"options"`
	Protocol        *string        `json:"protocol,omitempty"`
	SelectionFields OVSSet[string] `json:"selection_fields"`
	Vips            OVSMap[string] `json:"vips"`
}

// Load_Balancer_Group
type LoadBalancerGroupNB struct {
	LoadBalancer OVSSet[UUID] `json:"load_balancer"` //  Load_Balancer
	Name         *string      `json:"name"`
}

// Load_Balancer_Health_Check
type LoadBalancerHealthCheckNB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Options     OVSMap[string] `json:"options"`
	Vip         *string        `json:"vip"`
}

// Logical_Router
type LogicalRouterNB struct {
	Copp              *UUID          `json:"copp,omitempty"` //  Copp
	Enabled           *bool          `json:"enabled,omitempty"`
	ExternalIds       OVSMap[string] `json:"external_ids"`
	LoadBalancer      OVSSet[UUID]   `json:"load_balancer"`       //  Load_Balancer
	LoadBalancerGroup OVSSet[UUID]   `json:"load_balancer_group"` //  Load_Balancer_Group
	Name              *string        `json:"name"`
	Nat               OVSSet[UUID]   `json:"nat"` //  NAT
	Options           OVSMap[string] `json:"options"`
	Policies          OVSSet[UUID]   `json:"policies"`      //  Logical_Router_Policy
	Ports             OVSSet[UUID]   `json:"ports"`         //  Logical_Router_Port
	StaticRoutes      OVSSet[UUID]   `json:"static_routes"` //  Logical_Router_Static_Route
}

// Logical_Router_Policy
type LogicalRouterPolicyNB struct {
	Action      *string        `json:"action"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Match       *string        `json:"match"`
	Nexthop     *string        `json:"nexthop,omitempty"`
	Nexthops    OVSSet[string] `json:"nexthops"`
	Options     OVSMap[string] `json:"options"`
	Priority    *int           `json:"priority"`
}

// Logical_Router_Port
type LogicalRouterPortNB struct {
	Enabled        *bool          `json:"enabled,omitempty"`
	ExternalIds    OVSMap[string] `json:"external_ids"`
	GatewayChassis OVSSet[UUID]   `json:"gateway_chassis"`            //  Gateway_Chassis
	HaChassisGroup *UUID          `json:"ha_chassis_group,omitempty"` //  HA_Chassis_Group
	Ipv6Prefix     OVSSet[string] `json:"ipv6_prefix"`
	Ipv6RaConfigs  OVSMap[string] `json:"ipv6_ra_configs"`
	Mac            *string        `json:"mac"`
	Name           *string        `json:"name"`
	Networks       OVSSet[string] `json:"networks"`
	Options        OVSMap[string] `json:"options"`
	Peer           *string        `json:"peer,omitempty"`
}

// Logical_Router_Static_Route
type LogicalRouterStaticRouteNB struct {
	Bfd         *UUID          `json:"bfd,omitempty"` //  BFD
	ExternalIds OVSMap[string] `json:"external_ids"`
	IpPrefix    *string        `json:"ip_prefix"`
	Nexthop     *string        `json:"nexthop"`
	Options     OVSMap[string] `json:"options"`
	OutputPort  *string        `json:"output_port,omitempty"`
	Policy      *string        `json:"policy,omitempty"`
	RouteTable  *string        `json:"route_table"`
}

// Logical_Switch
type LogicalSwitchNB struct {
	Acls              OVSSet[UUID]   `json:"acls"`           //  ACL
	Copp              *UUID          `json:"copp,omitempty"` //  Copp
	DnsRecords        OVSSet[UUID]   `json:"dns_records"`    //  DNS
	ExternalIds       OVSMap[string] `json:"external_ids"`
	ForwardingGroups  OVSSet[UUID]   `json:"forwarding_groups"`   //  Forwarding_Group
	LoadBalancer      OVSSet[UUID]   `json:"load_balancer"`       //  Load_Balancer
	LoadBalancerGroup OVSSet[UUID]   `json:"load_balancer_group"` //  Load_Balancer_Group
	Name              *string        `json:"name"`
	OtherConfig       OVSMap[string] `json:"other_config"`
	Ports             OVSSet[UUID]   `json:"ports"`     //  Logical_Switch_Port
	QosRules          OVSSet[UUID]   `json:"qos_rules"` //  QoS
}

// Logical_Switch_Port
type LogicalSwitchPortNB struct {
	Addresses        OVSSet[string] `json:"addresses"`
	Dhcpv4Options    *UUID          `json:"dhcpv4_options,omitempty"` //  DHCP_Options
	Dhcpv6Options    *UUID          `json:"dhcpv6_options,omitempty"` //  DHCP_Options
	DynamicAddresses *string        `json:"dynamic_addresses,omitempty"`
	Enabled          *bool          `json:"enabled,omitempty"`
	ExternalIds      OVSMap[string] `json:"external_ids"`
	HaChassisGroup   *UUID          `json:"ha_chassis_group,omitempty"` //  HA_Chassis_Group
	Name             *string        `json:"name"`
	Options          OVSMap[string] `json:"options"`
	ParentName       *string        `json:"parent_name,omitempty"`
	PortSecurity     OVSSet[string] `json:"port_security"`
	Tag              *int           `json:"tag,omitempty"`
	TagRequest       *int           `json:"tag_request,omitempty"`
	Type             *string        `json:"type"`
	Up               *bool          `json:"up,omitempty"`
}

// Meter
type MeterNB struct {
	Bands       OVSSet[UUID]   `json:"bands"` //  Meter_Band
	ExternalIds OVSMap[string] `json:"external_ids"`
	Fair        *bool          `json:"fair,omitempty"`
	Name        *string        `json:"name"`
	Unit        *string        `json:"unit"`
}

// Meter_Band
type MeterBandNB struct {
	Action      *string        `json:"action"`
	BurstSize   *int           `json:"burst_size"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Rate        *int           `json:"rate"`
}

// NAT
type NATNB struct {
	AllowedExtIps     *UUID          `json:"allowed_ext_ips,omitempty"`  //  Address_Set
	ExemptedExtIps    *UUID          `json:"exempted_ext_ips,omitempty"` //  Address_Set
	ExternalIds       OVSMap[string] `json:"external_ids"`
	ExternalIp        *string        `json:"external_ip"`
	ExternalMac       *string        `json:"external_mac,omitempty"`
	ExternalPortRange *string        `json:"external_port_range"`
	LogicalIp         *string        `json:"logical_ip"`
	LogicalPort       *string        `json:"logical_port,omitempty"`
	Options           OVSMap[string] `json:"options"`
	Type              *string        `json:"type"`
}

// NB_Global
type NBGlobalNB struct {
	Connections    OVSSet[UUID]   `json:"connections"` //  Connection
	ExternalIds    OVSMap[string] `json:"external_ids"`
	HvCfg          *int           `json:"hv_cfg"`
	HvCfgTimestamp *int           `json:"hv_cfg_timestamp"`
	Ipsec          *bool          `json:"ipsec"`
	Name           *string        `json:"name"`
	NbCfg          *int           `json:"nb_cfg"`
	NbCfgTimestamp *int           `json:"nb_cfg_timestamp"`
	Options        OVSMap[string] `json:"options"`
	SbCfg          *int           `json:"sb_cfg"`
	SbCfgTimestamp *int           `json:"sb_cfg_timestamp"`
	Ssl            *UUID          `json:"ssl,omitempty"` //  SSL
}

// Port_Group
type PortGroupNB struct {
	Acls        OVSSet[UUID]   `json:"acls"` //  ACL
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        *string        `json:"name"`
	Ports       OVSSet[UUID]   `json:"ports"` //  Logical_Switch_Port
}

// QoS
type QoSNB struct {
	Action      OVSMap[int]    `json:"action"`
	Bandwidth   OVSMap[int]    `json:"bandwidth"`
	Direction   *string        `json:"direction"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Match       *string        `json:"match"`
	Priority    *int           `json:"priority"`
}

// SSL
type SSLNB struct {
	BootstrapCaCert *bool          `json:"bootstrap_ca_cert"`
	CaCert          *string        `json:"ca_cert"`
	Certificate     *string        `json:"certificate"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	PrivateKey      *string        `json:"private_key"`
	SslCiphers      *string        `json:"ssl_ciphers"`
	SslProtocols    *string        `json:"ssl_protocols"`
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
	return len(nb.NBGlobal) > 0
}
