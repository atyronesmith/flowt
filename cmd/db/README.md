# db

## Introduction

Generate information about an OVN database.

- Database schema in visual format (dot->svg)
- Pretty JSON data file
- Statistics about the database

## Usage

To get help...

`go run cmd/db/main.go -h`

Example run...

`go run cmd/db/main.go -v -o demo -jd -ds demo/ovnsb_db.db`

```text
Date            54479-03-08 13:00:36 +0000 UTC                         
Comment         ovn-northd                       
IsDiff          true                         
AddressSet      9
BFD             0
Chassis         2            
                         sos-novacompute-0.localdomain            
                         controller-0.localdomain
ChassisPrivate  2
Connection      1
ControllerEvent 0
DHCPOptions     37
DHCPv6Options   4
DNS             0
DatapathBinding 6
Encap           2
FDB             0
GatewayChassis  0
HAChassis       1
HAChassisGroup  1
IGMPGroup       0
IPMulticast     5
LoadBalancer    0
LogicalDPGroup  4
LogicalFlow     394
MACBinding      1
Meter           0
MeterBand       0
MulticastGroup  18
PortBinding     20
PortGroup       9
RBACPermission  9
RBACRole        1
SBGlobal        1
SSL             0
ServiceMonitor  0
Writing ovn_southbound.dot...
Writing ovn_southbound.json...

or

`go run cmd/db/main.go -v -o demo -jd -ds demo/ovnnb_db.db`

```text
Date                     54479-03-08 13:00:30 +0000 UTC   
Comment                  ""
IsDiff                   true   
ACL                      15
AddressSet               0
BFD                      0                      
Connection               1               
Copp                     0                     
DHCPOptions              2              
DNS                      0                      
ForwardingGroup          0          
GatewayChassis           1           
HAChassis                1                
HAChassisGroup           1           
LoadBalancer             0             
LoadBalancerGroup        0        
LoadBalancerHealthCheck  0  
LogicalRouter            1            
LogicalRouterPolicy      0      
LogicalRouterPort        2        
LogicalRouterStaticRoute 2 
LogicalSwitch            5            
                         neutron-d8953248-ba41-4ef4-b7a3-471afed8fd8f    3 Ports            
                         neutron-a6e858b0-c295-41d4-8ff4-858c18695d0c    4 Ports            
                         neutron-82e03259-6310-4b9d-9575-7f07d613ce09    4 Ports            
                             3 Ports            
                         neutron-561990d3-f4d5-431d-ae92-a85b83f4f570    3 Ports
LogicalSwitchPort        17 7 VMs, 2 Routers, 5 LPorts, 3 LNets
Meter                    0                    
MeterBand                0                
NAT                      2
                         dnat_and_snat    1
                                  snat    1                
Global                   1                 
PortGroup                4                
QoS                      0                      
SSL                      0
Writing ovn_northbound.dot...
Writing ovn_northbound.json...
