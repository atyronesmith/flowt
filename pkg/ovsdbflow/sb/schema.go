package sb

import (
	types "github.com/atyronesmith/flowt/pkg/ovsdb"
)

type AddressSet struct {
	Addresses types.OVSSet[string] `json:"addresses"`
	Name      string               `json:"name"`
}
type BFD struct {
	DetectMult  int                  `json:"detect_mult"`
	Disc        int                  `json:"disc"`
	DstIp       string               `json:"dst_ip"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	LogicalPort string               `json:"logical_port"`
	MinRx       int                  `json:"min_rx"`
	MinTx       int                  `json:"min_tx"`
	Options     types.OVSMap[string] `json:"options"`
	SrcPort     int                  `json:"src_port"`
	Status      string               `json:"status"`
}
type Chassis struct {
	Encaps              types.OVSSet[string] `json:"encaps"`
	ExternalIds         types.OVSMap[string] `json:"external_ids"`
	Hostname            string               `json:"hostname"`
	Name                string               `json:"name"`
	NbCfg               int                  `json:"nb_cfg"`
	OtherConfig         types.OVSMap[string] `json:"other_config"`
	TransportZones      types.OVSSet[string] `json:"transport_zones"`
	VtepLogicalSwitches types.OVSSet[string] `json:"vtep_logical_switches"`
}
type ChassisPrivate struct {
	Chassis        types.OVSSet[string] `json:"chassis"`
	ExternalIds    types.OVSMap[string] `json:"external_ids"`
	Name           string               `json:"name"`
	NbCfg          int                  `json:"nb_cfg"`
	NbCfgTimestamp int                  `json:"nb_cfg_timestamp"`
}
type Connection struct {
	ExternalIds     types.OVSMap[string] `json:"external_ids"`
	InactivityProbe int                  `json:"inactivity_probe"`
	IsConnected     bool                 `json:"is_connected"`
	MaxBackoff      int                  `json:"max_backoff"`
	OtherConfig     types.OVSMap[string] `json:"other_config"`
	ReadOnly        bool                 `json:"read_only"`
	Role            string               `json:"role"`
	Status          types.OVSMap[string] `json:"status"`
	Target          string               `json:"target"`
}
type ControllerEvent struct {
	Chassis   types.OVSSet[string] `json:"chassis"`
	EventInfo types.OVSMap[string] `json:"event_info"`
	EventType string               `json:"event_type"`
	SeqNum    int                  `json:"seq_num"`
}
type DHCPOptions struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type DHCPv6Options struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type DNS struct {
	Datapaths   types.OVSSet[string] `json:"datapaths"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Records     types.OVSMap[string] `json:"records"`
}
type DatapathBinding struct {
	ExternalIds   types.OVSMap[string] `json:"external_ids"`
	LoadBalancers types.OVSSet[string] `json:"load_balancers"`
	TunnelKey     int                  `json:"tunnel_key"`
}
type Encap struct {
	ChassisName string               `json:"chassis_name"`
	Ip          string               `json:"ip"`
	Options     types.OVSMap[string] `json:"options"`
	Type        string               `json:"type"`
}
type FDB struct {
	DpKey   int    `json:"dp_key"`
	Mac     string `json:"mac"`
	PortKey int    `json:"port_key"`
}
type GatewayChassis struct {
	Chassis     types.OVSSet[string] `json:"chassis"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Name        string               `json:"name"`
	Options     types.OVSMap[string] `json:"options"`
	Priority    int                  `json:"priority"`
}
type HAChassis struct {
	Chassis     types.OVSSet[string] `json:"chassis"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Priority    int                  `json:"priority"`
}
type HAChassisGroup struct {
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	HaChassis   types.OVSSet[string] `json:"ha_chassis"`
	Name        string               `json:"name"`
	RefChassis  types.OVSSet[string] `json:"ref_chassis"`
}
type IGMPGroup struct {
	Address  string               `json:"address"`
	Chassis  types.OVSSet[string] `json:"chassis"`
	Datapath types.OVSSet[string] `json:"datapath"`
	Ports    types.OVSSet[string] `json:"ports"`
}
type IPMulticast struct {
	Datapath      types.OVSSet[string] `json:"datapath"`
	Enabled       bool                 `json:"enabled"`
	EthSrc        string               `json:"eth_src"`
	IdleTimeout   int                  `json:"idle_timeout"`
	Ip4Src        string               `json:"ip4_src"`
	Ip6Src        string               `json:"ip6_src"`
	Querier       bool                 `json:"querier"`
	QueryInterval int                  `json:"query_interval"`
	QueryMaxResp  int                  `json:"query_max_resp"`
	SeqNo         int                  `json:"seq_no"`
	TableSize     int                  `json:"table_size"`
}
type LoadBalancer struct {
	Datapaths   types.OVSSet[string] `json:"datapaths"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Name        string               `json:"name"`
	Options     types.OVSMap[string] `json:"options"`
	Protocol    string               `json:"protocol"`
	Vips        types.OVSMap[string] `json:"vips"`
}
type LogicalDPGroup struct {
	Datapaths types.OVSSet[string] `json:"datapaths"`
}
type LogicalFlow struct {
	Actions         string               `json:"actions"`
	ControllerMeter string               `json:"controller_meter"`
	ExternalIds     types.OVSMap[string] `json:"external_ids"`
	LogicalDatapath types.OVSSet[string] `json:"logical_datapath"`
	LogicalDpGroup  types.OVSSet[string] `json:"logical_dp_group"`
	Match           string               `json:"match"`
	Pipeline        string               `json:"pipeline"`
	Priority        int                  `json:"priority"`
	TableId         int                  `json:"table_id"`
	Tags            types.OVSMap[string] `json:"tags"`
}
type MACBinding struct {
	Datapath    types.OVSSet[string] `json:"datapath"`
	Ip          string               `json:"ip"`
	LogicalPort string               `json:"logical_port"`
	Mac         string               `json:"mac"`
}
type Meter struct {
	Bands types.OVSSet[string] `json:"bands"`
	Name  string               `json:"name"`
	Unit  string               `json:"unit"`
}
type MeterBand struct {
	Action    string `json:"action"`
	BurstSize int    `json:"burst_size"`
	Rate      int    `json:"rate"`
}
type MulticastGroup struct {
	Datapath  types.OVSSet[string] `json:"datapath"`
	Name      string               `json:"name"`
	Ports     types.OVSSet[string] `json:"ports"`
	TunnelKey int                  `json:"tunnel_key"`
}
type PortBinding struct {
	Chassis          types.OVSSet[string] `json:"chassis"`
	Datapath         types.OVSSet[string] `json:"datapath"`
	Encap            types.OVSSet[string] `json:"encap"`
	ExternalIds      types.OVSMap[string] `json:"external_ids"`
	GatewayChassis   types.OVSSet[string] `json:"gateway_chassis"`
	HaChassisGroup   types.OVSSet[string] `json:"ha_chassis_group"`
	LogicalPort      string               `json:"logical_port"`
	Mac              types.OVSSet[string] `json:"mac"`
	NatAddresses     types.OVSSet[string] `json:"nat_addresses"`
	Options          types.OVSMap[string] `json:"options"`
	ParentPort       string               `json:"parent_port"`
	RequestedChassis types.OVSSet[string] `json:"requested_chassis"`
	Tag              int                  `json:"tag"`
	TunnelKey        int                  `json:"tunnel_key"`
	Type             string               `json:"type"`
	Up               bool                 `json:"up"`
	VirtualParent    string               `json:"virtual_parent"`
}
type PortGroup struct {
	Name  string               `json:"name"`
	Ports types.OVSSet[string] `json:"ports"`
}
type RBACPermission struct {
	Authorization types.OVSSet[string] `json:"authorization"`
	InsertDelete  bool                 `json:"insert_delete"`
	Table         string               `json:"table"`
	Update        types.OVSSet[string] `json:"update"`
}
type RBACRole struct {
	Name        string               `json:"name"`
	Permissions types.OVSSet[string] `json:"permissions"`
}
type SBGlobal struct {
	Connections types.OVSSet[string] `json:"connections"`
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Ipsec       bool                 `json:"ipsec"`
	NbCfg       int                  `json:"nb_cfg"`
	Options     types.OVSMap[string] `json:"options"`
	Ssl         types.OVSSet[string] `json:"ssl"`
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
type ServiceMonitor struct {
	ExternalIds types.OVSMap[string] `json:"external_ids"`
	Ip          string               `json:"ip"`
	LogicalPort string               `json:"logical_port"`
	Options     types.OVSMap[string] `json:"options"`
	Port        int                  `json:"port"`
	Protocol    string               `json:"protocol"`
	SrcIp       string               `json:"src_ip"`
	SrcMac      string               `json:"src_mac"`
	Status      string               `json:"status"`
}

type OVSdb struct {
	AddressSet      map[string]AddressSet      `json:"Address_Set"`
	BFD             map[string]BFD             `json:"BFD"`
	Chassis         map[string]Chassis         `json:"Chassis"`
	ChassisPrivate  map[string]ChassisPrivate  `json:"Chassis_Private"`
	Connection      map[string]Connection      `json:"Connection"`
	ControllerEvent map[string]ControllerEvent `json:"Controller_Event"`
	DHCPOptions     map[string]DHCPOptions     `json:"DHCP_Options"`
	DHCPv6Options   map[string]DHCPv6Options   `json:"DHCPv6_Options"`
	DNS             map[string]DNS             `json:"DNS"`
	DatapathBinding map[string]DatapathBinding `json:"Datapath_Binding"`
	Encap           map[string]Encap           `json:"Encap"`
	FDB             map[string]FDB             `json:"FDB"`
	GatewayChassis  map[string]GatewayChassis  `json:"Gateway_Chassis"`
	HAChassis       map[string]HAChassis       `json:"HA_Chassis"`
	HAChassisGroup  map[string]HAChassisGroup  `json:"HA_Chassis_Group"`
	IGMPGroup       map[string]IGMPGroup       `json:"IGMP_Group"`
	IPMulticast     map[string]IPMulticast     `json:"IP_Multicast"`
	LoadBalancer    map[string]LoadBalancer    `json:"Load_Balancer"`
	LogicalDPGroup  map[string]LogicalDPGroup  `json:"Logical_DP_Group"`
	LogicalFlow     map[string]LogicalFlow     `json:"Logical_Flow"`
	MACBinding      map[string]MACBinding      `json:"MAC_Binding"`
	Meter           map[string]Meter           `json:"Meter"`
	MeterBand       map[string]MeterBand       `json:"Meter_Band"`
	MulticastGroup  map[string]MulticastGroup  `json:"Multicast_Group"`
	PortBinding     map[string]PortBinding     `json:"Port_Binding"`
	PortGroup       map[string]PortGroup       `json:"Port_Group"`
	RBACPermission  map[string]RBACPermission  `json:"RBAC_Permission"`
	RBACRole        map[string]RBACRole        `json:"RBAC_Role"`
	SBGlobal        map[string]SBGlobal        `json:"SB_Global"`
	SSL             map[string]SSL             `json:"SSL"`
	ServiceMonitor  map[string]ServiceMonitor  `json:"Service_Monitor"`
}
