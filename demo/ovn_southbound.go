package dbtypes

// Type: OVN_Southbound
// Version: 20.21.0
// Tables: 31

// Address_Set
type AddressSetSB struct {
	Addresses OVSSet[string] `json:"addresses"`
	Name      *string        `json:"name"`
}

// BFD
type BFDSB struct {
	DetectMult  *int           `json:"detect_mult"`
	Disc        *int           `json:"disc"`
	DstIp       *string        `json:"dst_ip"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	LogicalPort *string        `json:"logical_port"`
	MinRx       *int           `json:"min_rx"`
	MinTx       *int           `json:"min_tx"`
	Options     OVSMap[string] `json:"options"`
	SrcPort     *int           `json:"src_port"`
	Status      *string        `json:"status"`
}

// Chassis
type ChassisSB struct {
	Encaps              OVSSet[UUID]   `json:"encaps"` //  Encap
	ExternalIds         OVSMap[string] `json:"external_ids"`
	Hostname            *string        `json:"hostname"`
	Name                *string        `json:"name"`
	NbCfg               *int           `json:"nb_cfg"`
	OtherConfig         OVSMap[string] `json:"other_config"`
	TransportZones      OVSSet[string] `json:"transport_zones"`
	VtepLogicalSwitches OVSSet[string] `json:"vtep_logical_switches"`
}

// Chassis_Private
type ChassisPrivateSB struct {
	Chassis        *UUID          `json:"chassis,omitempty"` //  Chassis
	ExternalIds    OVSMap[string] `json:"external_ids"`
	Name           *string        `json:"name"`
	NbCfg          *int           `json:"nb_cfg"`
	NbCfgTimestamp *int           `json:"nb_cfg_timestamp"`
}

// Connection
type ConnectionSB struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	InactivityProbe *int           `json:"inactivity_probe,omitempty"`
	IsConnected     *bool          `json:"is_connected"`
	MaxBackoff      *int           `json:"max_backoff,omitempty"`
	OtherConfig     OVSMap[string] `json:"other_config"`
	ReadOnly        *bool          `json:"read_only"`
	Role            *string        `json:"role"`
	Status          OVSMap[string] `json:"status"`
	Target          *string        `json:"target"`
}

// Controller_Event
type ControllerEventSB struct {
	Chassis   *UUID          `json:"chassis,omitempty"` //  Chassis
	EventInfo OVSMap[string] `json:"event_info"`
	EventType *string        `json:"event_type"`
	SeqNum    *int           `json:"seq_num"`
}

// DHCP_Options
type DHCPOptionsSB struct {
	Code *int    `json:"code"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

// DHCPv6_Options
type DHCPv6OptionsSB struct {
	Code *int    `json:"code"`
	Name *string `json:"name"`
	Type *string `json:"type"`
}

// DNS
type DNSSB struct {
	Datapaths   OVSSet[UUID]   `json:"datapaths"` //  Datapath_Binding
	ExternalIds OVSMap[string] `json:"external_ids"`
	Records     OVSMap[string] `json:"records"`
}

// Datapath_Binding
type DatapathBindingSB struct {
	ExternalIds   OVSMap[string] `json:"external_ids"`
	LoadBalancers OVSSet[UUID]   `json:"load_balancers"`
	TunnelKey     *int           `json:"tunnel_key"`
}

// Encap
type EncapSB struct {
	ChassisName *string        `json:"chassis_name"`
	Ip          *string        `json:"ip"`
	Options     OVSMap[string] `json:"options"`
	Type        *string        `json:"type"`
}

// FDB
type FDBSB struct {
	DpKey   *int    `json:"dp_key"`
	Mac     *string `json:"mac"`
	PortKey *int    `json:"port_key"`
}

// Gateway_Chassis
type GatewayChassisSB struct {
	Chassis     *UUID          `json:"chassis,omitempty"` //  Chassis
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        *string        `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Priority    *int           `json:"priority"`
}

// HA_Chassis
type HAChassisSB struct {
	Chassis     *UUID          `json:"chassis,omitempty"` //  Chassis
	ExternalIds OVSMap[string] `json:"external_ids"`
	Priority    *int           `json:"priority"`
}

// HA_Chassis_Group
type HAChassisGroupSB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	HaChassis   OVSSet[UUID]   `json:"ha_chassis"` //  HA_Chassis
	Name        *string        `json:"name"`
	RefChassis  OVSSet[UUID]   `json:"ref_chassis"` //  Chassis
}

// IGMP_Group
type IGMPGroupSB struct {
	Address  *string      `json:"address"`
	Chassis  *UUID        `json:"chassis,omitempty"`  //  Chassis
	Datapath *UUID        `json:"datapath,omitempty"` //  Datapath_Binding
	Ports    OVSSet[UUID] `json:"ports"`              //  Port_Binding
}

// IP_Multicast
type IPMulticastSB struct {
	Datapath      *UUID   `json:"datapath"` //  Datapath_Binding
	Enabled       *bool   `json:"enabled,omitempty"`
	EthSrc        *string `json:"eth_src"`
	IdleTimeout   *int    `json:"idle_timeout,omitempty"`
	Ip4Src        *string `json:"ip4_src"`
	Ip6Src        *string `json:"ip6_src"`
	Querier       *bool   `json:"querier,omitempty"`
	QueryInterval *int    `json:"query_interval,omitempty"`
	QueryMaxResp  *int    `json:"query_max_resp,omitempty"`
	SeqNo         *int    `json:"seq_no"`
	TableSize     *int    `json:"table_size,omitempty"`
}

// Load_Balancer
type LoadBalancerSB struct {
	Datapaths   OVSSet[UUID]   `json:"datapaths"` //  Datapath_Binding
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        *string        `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Protocol    *string        `json:"protocol,omitempty"`
	Vips        OVSMap[string] `json:"vips"`
}

// Logical_DP_Group
type LogicalDPGroupSB struct {
	Datapaths OVSSet[UUID] `json:"datapaths"` //  Datapath_Binding
}

// Logical_Flow
type LogicalFlowSB struct {
	Actions         *string        `json:"actions"`
	ControllerMeter *string        `json:"controller_meter,omitempty"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	LogicalDatapath *UUID          `json:"logical_datapath,omitempty"` //  Datapath_Binding
	LogicalDpGroup  *UUID          `json:"logical_dp_group,omitempty"` //  Logical_DP_Group
	Match           *string        `json:"match"`
	Pipeline        *string        `json:"pipeline"`
	Priority        *int           `json:"priority"`
	TableId         *int           `json:"table_id"`
	Tags            OVSMap[string] `json:"tags"`
}

// MAC_Binding
type MACBindingSB struct {
	Datapath    *UUID   `json:"datapath"` //  Datapath_Binding
	Ip          *string `json:"ip"`
	LogicalPort *string `json:"logical_port"`
	Mac         *string `json:"mac"`
}

// Meter
type MeterSB struct {
	Bands OVSSet[UUID] `json:"bands"` //  Meter_Band
	Name  *string      `json:"name"`
	Unit  *string      `json:"unit"`
}

// Meter_Band
type MeterBandSB struct {
	Action    *string `json:"action"`
	BurstSize *int    `json:"burst_size"`
	Rate      *int    `json:"rate"`
}

// Multicast_Group
type MulticastGroupSB struct {
	Datapath  *UUID        `json:"datapath"` //  Datapath_Binding
	Name      *string      `json:"name"`
	Ports     OVSSet[UUID] `json:"ports"` //  Port_Binding
	TunnelKey *int         `json:"tunnel_key"`
}

// Port_Binding
type PortBindingSB struct {
	Chassis          *UUID          `json:"chassis,omitempty"` //  Chassis
	Datapath         *UUID          `json:"datapath"`          //  Datapath_Binding
	Encap            *UUID          `json:"encap,omitempty"`   //  Encap
	ExternalIds      OVSMap[string] `json:"external_ids"`
	GatewayChassis   OVSSet[UUID]   `json:"gateway_chassis"`            //  Gateway_Chassis
	HaChassisGroup   *UUID          `json:"ha_chassis_group,omitempty"` //  HA_Chassis_Group
	LogicalPort      *string        `json:"logical_port"`
	Mac              OVSSet[string] `json:"mac"`
	NatAddresses     OVSSet[string] `json:"nat_addresses"`
	Options          OVSMap[string] `json:"options"`
	ParentPort       *string        `json:"parent_port,omitempty"`
	RequestedChassis *UUID          `json:"requested_chassis,omitempty"` //  Chassis
	Tag              *int           `json:"tag,omitempty"`
	TunnelKey        *int           `json:"tunnel_key"`
	Type             *string        `json:"type"`
	Up               *bool          `json:"up,omitempty"`
	VirtualParent    *string        `json:"virtual_parent,omitempty"`
}

// Port_Group
type PortGroupSB struct {
	Name  *string        `json:"name"`
	Ports OVSSet[string] `json:"ports"`
}

// RBAC_Permission
type RBACPermissionSB struct {
	Authorization OVSSet[string] `json:"authorization"`
	InsertDelete  *bool          `json:"insert_delete"`
	Table         *string        `json:"table"`
	Update        OVSSet[string] `json:"update"`
}

// RBAC_Role
type RBACRoleSB struct {
	Name        *string      `json:"name"`
	Permissions OVSMap[UUID] `json:"permissions"` //  RBAC_Permission
}

// SB_Global
type SBGlobalSB struct {
	Connections OVSSet[UUID]   `json:"connections"` //  Connection
	ExternalIds OVSMap[string] `json:"external_ids"`
	Ipsec       *bool          `json:"ipsec"`
	NbCfg       *int           `json:"nb_cfg"`
	Options     OVSMap[string] `json:"options"`
	Ssl         *UUID          `json:"ssl,omitempty"` //  SSL
}

// SSL
type SSLSB struct {
	BootstrapCaCert *bool          `json:"bootstrap_ca_cert"`
	CaCert          *string        `json:"ca_cert"`
	Certificate     *string        `json:"certificate"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	PrivateKey      *string        `json:"private_key"`
	SslCiphers      *string        `json:"ssl_ciphers"`
	SslProtocols    *string        `json:"ssl_protocols"`
}

// Service_Monitor
type ServiceMonitorSB struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Ip          *string        `json:"ip"`
	LogicalPort *string        `json:"logical_port"`
	Options     OVSMap[string] `json:"options"`
	Port        *int           `json:"port"`
	Protocol    *string        `json:"protocol,omitempty"`
	SrcIp       *string        `json:"src_ip"`
	SrcMac      *string        `json:"src_mac"`
	Status      *string        `json:"status,omitempty"`
}

type OVNSouthbound struct {
	Date            Time                         `json:"_date"`
	Comment         string                       `json:"_comment"`
	IsDiff          bool                         `json:"_is_diff"`
	AddressSet      map[string]AddressSetSB      `json:"Address_Set"`
	BFD             map[string]BFDSB             `json:"BFD"`
	Chassis         map[string]ChassisSB         `json:"Chassis"`
	ChassisPrivate  map[string]ChassisPrivateSB  `json:"Chassis_Private"`
	Connection      map[string]ConnectionSB      `json:"Connection"`
	ControllerEvent map[string]ControllerEventSB `json:"Controller_Event"`
	DHCPOptions     map[string]DHCPOptionsSB     `json:"DHCP_Options"`
	DHCPv6Options   map[string]DHCPv6OptionsSB   `json:"DHCPv6_Options"`
	DNS             map[string]DNSSB             `json:"DNS"`
	DatapathBinding map[string]DatapathBindingSB `json:"Datapath_Binding"`
	Encap           map[string]EncapSB           `json:"Encap"`
	FDB             map[string]FDBSB             `json:"FDB"`
	GatewayChassis  map[string]GatewayChassisSB  `json:"Gateway_Chassis"`
	HAChassis       map[string]HAChassisSB       `json:"HA_Chassis"`
	HAChassisGroup  map[string]HAChassisGroupSB  `json:"HA_Chassis_Group"`
	IGMPGroup       map[string]IGMPGroupSB       `json:"IGMP_Group"`
	IPMulticast     map[string]IPMulticastSB     `json:"IP_Multicast"`
	LoadBalancer    map[string]LoadBalancerSB    `json:"Load_Balancer"`
	LogicalDPGroup  map[string]LogicalDPGroupSB  `json:"Logical_DP_Group"`
	LogicalFlow     map[string]LogicalFlowSB     `json:"Logical_Flow"`
	MACBinding      map[string]MACBindingSB      `json:"MAC_Binding"`
	Meter           map[string]MeterSB           `json:"Meter"`
	MeterBand       map[string]MeterBandSB       `json:"Meter_Band"`
	MulticastGroup  map[string]MulticastGroupSB  `json:"Multicast_Group"`
	PortBinding     map[string]PortBindingSB     `json:"Port_Binding"`
	PortGroup       map[string]PortGroupSB       `json:"Port_Group"`
	RBACPermission  map[string]RBACPermissionSB  `json:"RBAC_Permission"`
	RBACRole        map[string]RBACRoleSB        `json:"RBAC_Role"`
	SBGlobal        map[string]SBGlobalSB        `json:"SB_Global"`
	SSL             map[string]SSLSB             `json:"SSL"`
	ServiceMonitor  map[string]ServiceMonitorSB  `json:"Service_Monitor"`
}

func (sb OVNSouthbound) IsValid() bool {
	return len(sb.SBGlobal) > 0
}
