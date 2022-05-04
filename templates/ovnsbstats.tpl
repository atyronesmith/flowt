Date            {{.Date}}                         
Comment         {{.Comment}}                       
IsDiff          {{.IsDiff}}                         
AddressSet      {{ len .AddressSet}}
BFD             {{ len .BFD}}
Chassis         {{ len .Chassis}}
{{- range .Chassis }}            
                         {{ .Hostname }}
{{- end }}
ChassisPrivate  {{ len .ChassisPrivate}}
Connection      {{ len .Connection}}
ControllerEvent {{ len .ControllerEvent}}
DHCPOptions     {{ len .DHCPOptions}}
DHCPv6Options   {{ len .DHCPv6Options}}
DNS             {{ len .DNS}}
DatapathBinding {{ len .DatapathBinding}}
Encap           {{ len .Encap}}
FDB             {{ len .FDB}}
GatewayChassis  {{ len .GatewayChassis}}
HAChassis       {{ len .HAChassis}}
HAChassisGroup  {{ len .HAChassisGroup}}
IGMPGroup       {{ len .IGMPGroup}}
IPMulticast     {{ len .IPMulticast}}
LoadBalancer    {{ len .LoadBalancer}}
LogicalDPGroup  {{ len .LogicalDPGroup}}
LogicalFlow     {{ len .LogicalFlow}}
MACBinding      {{ len .MACBinding}}
Meter           {{ len .Meter}}
MeterBand       {{ len .MeterBand}}
MulticastGroup  {{ len .MulticastGroup}}
PortBinding     {{ len .PortBinding}}
PortGroup       {{ len .PortGroup}}
RBACPermission  {{ len .RBACPermission}}
RBACRole        {{ len .RBACRole}}
SBGlobal        {{ len .SBGlobal}}
SSL             {{ len .SSL}}
ServiceMonitor  {{ len .ServiceMonitor}}
