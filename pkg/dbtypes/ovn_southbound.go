package dbtypes

// Type: OVN_Southbound
// Version: 20.21.0
// Tables: 31

// Address_Set
type SBAddressSet struct {
	Addresses OVSSet[string] `json:"addresses"`
	Name      string         `json:"name"`
}

// BFD
type SBBFD struct {
	DetectMult  int            `json:"detect_mult"`
	Disc        int            `json:"disc"`
	DstIp       string         `json:"dst_ip"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	LogicalPort string         `json:"logical_port"`
	MinRx       int            `json:"min_rx"`
	MinTx       int            `json:"min_tx"`
	Options     OVSMap[string] `json:"options"`
	SrcPort     int            `json:"src_port"`
	Status      string         `json:"status"`
}

// Chassis
type SBChassis struct {
	Encaps              OVSSet[UUID]   `json:"encaps"`
	ExternalIds         OVSMap[string] `json:"external_ids"`
	Hostname            string         `json:"hostname"`
	Name                string         `json:"name"`
	NbCfg               int            `json:"nb_cfg"`
	OtherConfig         OVSMap[string] `json:"other_config"`
	TransportZones      OVSSet[string] `json:"transport_zones"`
	VtepLogicalSwitches OVSSet[string] `json:"vtep_logical_switches"`
}

// Chassis_Private
type SBChassisPrivate struct {
	Chassis        OVSSet[UUID]   `json:"chassis"`
	ExternalIds    OVSMap[string] `json:"external_ids"`
	Name           string         `json:"name"`
	NbCfg          int            `json:"nb_cfg"`
	NbCfgTimestamp int            `json:"nb_cfg_timestamp"`
}

// Connection
type SBConnection struct {
	ExternalIds     OVSMap[string] `json:"external_ids"`
	InactivityProbe int            `json:"inactivity_probe"`
	IsConnected     bool           `json:"is_connected"`
	MaxBackoff      int            `json:"max_backoff"`
	OtherConfig     OVSMap[string] `json:"other_config"`
	ReadOnly        bool           `json:"read_only"`
	Role            string         `json:"role"`
	Status          OVSMap[string] `json:"status"`
	Target          string         `json:"target"`
}

// Controller_Event
type SBControllerEvent struct {
	Chassis   OVSSet[UUID]   `json:"chassis"`
	EventInfo OVSMap[string] `json:"event_info"`
	EventType string         `json:"event_type"`
	SeqNum    int            `json:"seq_num"`
}

// DHCP_Options
type SBDHCPOptions struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// DHCPv6_Options
type SBDHCPv6Options struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// DNS
type SBDNS struct {
	Datapaths   OVSSet[UUID]   `json:"datapaths"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Records     OVSMap[string] `json:"records"`
}

// Datapath_Binding
type SBDatapathBinding struct {
	ExternalIds   OVSMap[string] `json:"external_ids"`
	LoadBalancers OVSSet[UUID]   `json:"load_balancers"`
	TunnelKey     int            `json:"tunnel_key"`
}

// Encap
type SBEncap struct {
	ChassisName string         `json:"chassis_name"`
	Ip          string         `json:"ip"`
	Options     OVSMap[string] `json:"options"`
	Type        string         `json:"type"`
}

// FDB
type SBFDB struct {
	DpKey   int    `json:"dp_key"`
	Mac     string `json:"mac"`
	PortKey int    `json:"port_key"`
}

// Gateway_Chassis
type SBGatewayChassis struct {
	Chassis     OVSSet[UUID]   `json:"chassis"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        string         `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Priority    int            `json:"priority"`
}

// HA_Chassis
type SBHAChassis struct {
	Chassis     OVSSet[UUID]   `json:"chassis"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Priority    int            `json:"priority"`
}

// HA_Chassis_Group
type SBHAChassisGroup struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	HaChassis   OVSSet[UUID]   `json:"ha_chassis"`
	Name        string         `json:"name"`
	RefChassis  OVSSet[UUID]   `json:"ref_chassis"`
}

// IGMP_Group
type SBIGMPGroup struct {
	Address  string       `json:"address"`
	Chassis  OVSSet[UUID] `json:"chassis"`
	Datapath OVSSet[UUID] `json:"datapath"`
	Ports    OVSSet[UUID] `json:"ports"`
}

// IP_Multicast
type SBIPMulticast struct {
	Datapath      OVSSet[UUID] `json:"datapath"`
	Enabled       bool         `json:"enabled"`
	EthSrc        string       `json:"eth_src"`
	IdleTimeout   int          `json:"idle_timeout"`
	Ip4Src        string       `json:"ip4_src"`
	Ip6Src        string       `json:"ip6_src"`
	Querier       bool         `json:"querier"`
	QueryInterval int          `json:"query_interval"`
	QueryMaxResp  int          `json:"query_max_resp"`
	SeqNo         int          `json:"seq_no"`
	TableSize     int          `json:"table_size"`
}

// Load_Balancer
type SBLoadBalancer struct {
	Datapaths   OVSSet[UUID]   `json:"datapaths"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Name        string         `json:"name"`
	Options     OVSMap[string] `json:"options"`
	Protocol    string         `json:"protocol"`
	Vips        OVSMap[string] `json:"vips"`
}

// Logical_DP_Group
type SBLogicalDPGroup struct {
	Datapaths OVSSet[UUID] `json:"datapaths"`
}

// Logical_Flow
type SBLogicalFlow struct {
	Actions         string         `json:"actions"`
	ControllerMeter string         `json:"controller_meter"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	LogicalDatapath OVSSet[UUID]   `json:"logical_datapath"`
	LogicalDpGroup  OVSSet[UUID]   `json:"logical_dp_group"`
	Match           string         `json:"match"`
	Pipeline        string         `json:"pipeline"`
	Priority        int            `json:"priority"`
	TableId         int            `json:"table_id"`
	Tags            OVSMap[string] `json:"tags"`
}

// MAC_Binding
type SBMACBinding struct {
	Datapath    OVSSet[UUID] `json:"datapath"`
	Ip          string       `json:"ip"`
	LogicalPort string       `json:"logical_port"`
	Mac         string       `json:"mac"`
}

// Meter
type SBMeter struct {
	Bands OVSSet[UUID] `json:"bands"`
	Name  string       `json:"name"`
	Unit  string       `json:"unit"`
}

// Meter_Band
type SBMeterBand struct {
	Action    string `json:"action"`
	BurstSize int    `json:"burst_size"`
	Rate      int    `json:"rate"`
}

// Multicast_Group
type SBMulticastGroup struct {
	Datapath  OVSSet[UUID] `json:"datapath"`
	Name      string       `json:"name"`
	Ports     OVSSet[UUID] `json:"ports"`
	TunnelKey int          `json:"tunnel_key"`
}

// Port_Binding
type SBPortBinding struct {
	Chassis          OVSSet[UUID]   `json:"chassis"`
	Datapath         OVSSet[UUID]   `json:"datapath"`
	Encap            OVSSet[UUID]   `json:"encap"`
	ExternalIds      OVSMap[string] `json:"external_ids"`
	GatewayChassis   OVSSet[UUID]   `json:"gateway_chassis"`
	HaChassisGroup   OVSSet[UUID]   `json:"ha_chassis_group"`
	LogicalPort      string         `json:"logical_port"`
	Mac              OVSSet[string] `json:"mac"`
	NatAddresses     OVSSet[string] `json:"nat_addresses"`
	Options          OVSMap[string] `json:"options"`
	ParentPort       string         `json:"parent_port"`
	RequestedChassis OVSSet[UUID]   `json:"requested_chassis"`
	Tag              int            `json:"tag"`
	TunnelKey        int            `json:"tunnel_key"`
	Type             string         `json:"type"`
	Up               bool           `json:"up"`
	VirtualParent    string         `json:"virtual_parent"`
}

// Port_Group
type SBPortGroup struct {
	Name  string         `json:"name"`
	Ports OVSSet[string] `json:"ports"`
}

// RBAC_Permission
type SBRBACPermission struct {
	Authorization OVSSet[string] `json:"authorization"`
	InsertDelete  bool           `json:"insert_delete"`
	Table         string         `json:"table"`
	Update        OVSSet[string] `json:"update"`
}

// RBAC_Role
type SBRBACRole struct {
	Name        string       `json:"name"`
	Permissions OVSMap[UUID] `json:"permissions"`
}

// SB_Global
type SBSBGlobal struct {
	Connections OVSSet[UUID]   `json:"connections"`
	ExternalIds OVSMap[string] `json:"external_ids"`
	Ipsec       bool           `json:"ipsec"`
	NbCfg       int            `json:"nb_cfg"`
	Options     OVSMap[string] `json:"options"`
	Ssl         OVSSet[UUID]   `json:"ssl"`
}

// SSL
type SBSSL struct {
	BootstrapCaCert bool           `json:"bootstrap_ca_cert"`
	CaCert          string         `json:"ca_cert"`
	Certificate     string         `json:"certificate"`
	ExternalIds     OVSMap[string] `json:"external_ids"`
	PrivateKey      string         `json:"private_key"`
	SslCiphers      string         `json:"ssl_ciphers"`
	SslProtocols    string         `json:"ssl_protocols"`
}

// Service_Monitor
type SBServiceMonitor struct {
	ExternalIds OVSMap[string] `json:"external_ids"`
	Ip          string         `json:"ip"`
	LogicalPort string         `json:"logical_port"`
	Options     OVSMap[string] `json:"options"`
	Port        int            `json:"port"`
	Protocol    string         `json:"protocol"`
	SrcIp       string         `json:"src_ip"`
	SrcMac      string         `json:"src_mac"`
	Status      string         `json:"status"`
}

type OVNSouthbound struct {
	Date              Time                         `json:"_date"`
	Comment           string                       `json:"_comment"`
	IsDiff            bool                         `json:"_is_diff"`
	SBAddressSet      map[string]SBAddressSet      `json:"Address_Set"`
	SBBFD             map[string]SBBFD             `json:"BFD"`
	SBChassis         map[string]SBChassis         `json:"Chassis"`
	SBChassisPrivate  map[string]SBChassisPrivate  `json:"Chassis_Private"`
	SBConnection      map[string]SBConnection      `json:"Connection"`
	SBControllerEvent map[string]SBControllerEvent `json:"Controller_Event"`
	SBDHCPOptions     map[string]SBDHCPOptions     `json:"DHCP_Options"`
	SBDHCPv6Options   map[string]SBDHCPv6Options   `json:"DHCPv6_Options"`
	SBDNS             map[string]SBDNS             `json:"DNS"`
	SBDatapathBinding map[string]SBDatapathBinding `json:"Datapath_Binding"`
	SBEncap           map[string]SBEncap           `json:"Encap"`
	SBFDB             map[string]SBFDB             `json:"FDB"`
	SBGatewayChassis  map[string]SBGatewayChassis  `json:"Gateway_Chassis"`
	SBHAChassis       map[string]SBHAChassis       `json:"HA_Chassis"`
	SBHAChassisGroup  map[string]SBHAChassisGroup  `json:"HA_Chassis_Group"`
	SBIGMPGroup       map[string]SBIGMPGroup       `json:"IGMP_Group"`
	SBIPMulticast     map[string]SBIPMulticast     `json:"IP_Multicast"`
	SBLoadBalancer    map[string]SBLoadBalancer    `json:"Load_Balancer"`
	SBLogicalDPGroup  map[string]SBLogicalDPGroup  `json:"Logical_DP_Group"`
	SBLogicalFlow     map[string]SBLogicalFlow     `json:"Logical_Flow"`
	SBMACBinding      map[string]SBMACBinding      `json:"MAC_Binding"`
	SBMeter           map[string]SBMeter           `json:"Meter"`
	SBMeterBand       map[string]SBMeterBand       `json:"Meter_Band"`
	SBMulticastGroup  map[string]SBMulticastGroup  `json:"Multicast_Group"`
	SBPortBinding     map[string]SBPortBinding     `json:"Port_Binding"`
	SBPortGroup       map[string]SBPortGroup       `json:"Port_Group"`
	SBRBACPermission  map[string]SBRBACPermission  `json:"RBAC_Permission"`
	SBRBACRole        map[string]SBRBACRole        `json:"RBAC_Role"`
	SBSBGlobal        map[string]SBSBGlobal        `json:"SB_Global"`
	SBSSL             map[string]SBSSL             `json:"SSL"`
	SBServiceMonitor  map[string]SBServiceMonitor  `json:"Service_Monitor"`
}

func (nb *OVNSouthbound) IsValid() bool {
	return len(nb.SBLogicalFlow) > 0
}
