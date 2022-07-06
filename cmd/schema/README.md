# schema

Generate go source files that map OVN Northbound or Southbound objects into go structs.  The data structures are used in conjungtion with json unmarshalling tools to load the contents of either a southbound or nortbound OVN database.  The generated files can be placed in the pkg/dbtypes directory of this project.  

The **schema** program only generates the source files related to the schema.  The following files need to be included in the same package.

- ovs_map.go
- ovs_set.go
- unix_time.go
- uuid.go

These files are located in this project *at pkg/dbtypes/*.

## Example

`go run cmd/schema/main.go -o example/ovn_southbound.go example/ovnsb_db.db`

or

`go run cmd/schema/main.go -o example/ovn_northbound.go example/ovnnb_db.db`

An example generated code is below.  

```go
package dbtypes

// Type: OVN_Northbound
// Version: 5.34.1
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
 Priority    *int           `json:"priority"`
 Severity    *string        `json:"severity,omitempty"`
}
...
.skip.
...
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
...
```
