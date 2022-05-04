Date                     {{.Db.Date}}   
Comment                  "{{.Db.Comment}}"
IsDiff                   {{.Db.IsDiff}}   
ACL                      {{len .Db.ACL}}
AddressSet               {{len .Db.AddressSet}}
BFD                      {{len .Db.BFD}}                      
Connection               {{len .Db.Connection}}               
Copp                     {{len .Db.Copp}}                     
DHCPOptions              {{len .Db.DHCPOptions}}              
DNS                      {{len .Db.DNS}}                      
ForwardingGroup          {{len .Db.ForwardingGroup}}          
GatewayChassis           {{len .Db.GatewayChassis}}           
HAChassis                {{len .Db.HAChassis}}                
HAChassisGroup           {{len .Db.HAChassisGroup}}           
LoadBalancer             {{len .Db.LoadBalancer}}             
LoadBalancerGroup        {{len .Db.LoadBalancerGroup}}        
LoadBalancerHealthCheck  {{len .Db.LoadBalancerHealthCheck}}  
LogicalRouter            {{len .Db.LogicalRouter}}            
LogicalRouterPolicy      {{len .Db.LogicalRouterPolicy}}      
LogicalRouterPort        {{len .Db.LogicalRouterPort}}        
LogicalRouterStaticRoute {{len .Db.LogicalRouterStaticRoute}} 
LogicalSwitch            {{len .Db.LogicalSwitch}}
{{- range .Db.LogicalSwitch }}            
                         {{ .Name }} {{ $l := len .Ports }}{{ printf "%4d" $l }} Ports
{{- end }}
LogicalSwitchPort        {{len .Db.LogicalSwitchPort}} {{.Stats.NumVMPorts}} VMs, {{.Stats.NumRouterPorts}} Routers, {{.Stats.NumLocalPorts}} LPorts, {{.Stats.NumLocalNetPorts}} LNets
Meter                    {{len .Db.Meter}}                    
MeterBand                {{len .Db.MeterBand}}                
NAT                      {{len .Db.NAT}}      
{{- range $k, $v := .Stats.NAT }}
                         {{ printf "%13s %4d" $k $v }}
{{- end}}                
Global                   {{len .Db.NBGlobal}}                 
PortGroup                {{len .Db.PortGroup}}                
QoS                      {{len .Db.QoS}}                      
SSL                      {{len .Db.SSL}}
