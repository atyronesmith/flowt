#!/bin/bash
{{range $dhcpKey, $dhcp := .Db.DHCPOptions -}}
CIDR_{{ DashToUnder $dhcpKey}}=$(ovn-nbctl create dhcp_options cidr={{$dhcp.Cidr}} \
  options='{{BuildMap $dhcp.Options}}')
{{end -}}

{{- range $switchKey, $switch := .Db.LogicalSwitch}}
{{- $switchName := MapSetString $switch.Name $switchKey}}
ovn-nbctl ls-add {{ $switchName }}
{{- range $portRef := .Ports}}{{$port := LookupLogicalSwitchPort $.Db.LogicalSwitchPort $portRef}}
{{- if $port}}
ovn-nbctl lsp-add {{ $switchName }} {{ $port.Name}} 
ovn-nbctl lsp-set-addresses {{$port.Name}} {{ BuildAddresses $port.Addresses }}
{{- if $port.PortSecurity}}
ovn-nbctl lsp-set-port-security {{$port.Name}} {{ BuildAddresses $port.Addresses }}
{{- end}}
{{- if $port.Enabled}}
ovn-nbctl lsp-set-enabled {{$port.Name}} {{$port.Enabled}}
{{- end}}
{{- if $port.Type}}
ovn-nbctl lsp-set-type {{ $switchName }} {{$port.Name}} {{$port.Type}}
{{- end}}
{{- range $optionKey, $optionValue := $port.Options}}
ovn-nbctl lsp-set-options {{ $switchName }} {{$optionKey}}={{$optionValue}}
{{- end}}
{{- if $port.Dhcpv4Options}}
ovn-nbctl lsp-set-dhcp4-options "$CIDR_{{AccessStringSlice $port.Dhcpv4Options}}"
{{- end}}
{{- end}}
{{- end}}
{{- range $switch.Acls}} 
{{- $uuid := UUIDToString .}}
{{- $acl := index $.Db.ACL $uuid}}
{{- /* TODO -- --after-lb --meter */}}
ovn-nbctl acl-add --type=switch {{if $acl.Log}}--log {{end}}{{if $acl.Severity}}--severity={{$acl.Severity}} {{end}}{{if $acl.Name}}--name={{$acl.Name}} {{end}}{{if $acl.Label}}--label={{$acl.Label}} {{end}}{{$switchName}} {{$acl.Direction}} {{$acl.Priority}} '{{$acl.Match}}' {{$acl.Action}}
{{- end }}
{{end -}}

{{- range $routerKey, $router := .Db.LogicalRouter}}
{{- $routerName := MapSetString $router.Name $routerKey}}
ovn-nbctl lr-add {{ $routerName }}
{{- range $optionKey, $optionValue := $router.Options}}
ovn-nbctl set logical_router {{ $routerName }} options:{{$optionKey}}={{$optionValue}}
{{- end}}
{{- range $portRef := .Ports -}}
{{- $port := LookupLogicalRouterPort $.Db.LogicalRouterPort $portRef }}
{{- if $port }}
ovn-nbctl lrp-add {{$routerName}} {{$port.Name}} {{$port.Mac}} {{range $port.Networks}}{{.}}{{end}} {{if $port.Peer }}peer={{$port.Peer}}{{end}}
{{- if $port.Enabled -}}
ovn-nbctl lrp-set-enabled {{$port.Name}} {{$port.Enabled}}
{{- end -}}
{{- if $port.GatewayChassis}}
ovn-nbctl lrp-set-gateway-chassis {{$port.Name}} {{range $port.GatewayChassis}}{{.}}{{end}}
{{- end -}}
{{- end }} 
{{- end }}
{{- range $portRef := .StaticRoutes -}}
{{- $port := LookupLogicalRouterStaticRoute $.Db.LogicalRouterStaticRoute $portRef }}
{{- if $port }}
ovn-nbctl lr-route-add {{$routerName}} {{$port.IpPrefix}} {{$port.Nexthop}} {{if $port.OutputPort}}{{.}}{{end}}
{{- end }} 
{{- end }}
{{- range $natRef := .Nat -}}
{{- $nat := LookupNAT $.Db.NAT $natRef }}
{{ if $nat }}
{{- $isStateless := index $nat.Options "stateless" -}}
ovn-nbctl lr-nat-add{{if $isStateless}} --stateless{{end}} {{$routerName}} {{$nat.Type}} {{$nat.ExternalIp}} {{$nat.LogicalIp}} {{if and $nat.LogicalPort $nat.ExternalMac}}{{$nat.LogicalPort}} {{$nat.ExternalMac}}{{end}}
{{- end }}
{{- end }}
{{ end -}}

{{- range $pgKey, $pg := .Db.PortGroup}}
{{- $pgName := MapSetString $pg.Name $pgKey}}
ovn-nbctl pg-add {{$pgName}} 
{{- range $pg.Ports}} 
{{- $uuid := UUIDToString .}}
{{- $p := index $.Db.LogicalSwitchPort $uuid}}
ovn-nbctl pg-set-ports {{$pgName}} {{ $p.Name }}
{{- end }}
{{- range $pg.Acls}} 
{{- $uuid := UUIDToString .}}
{{- $acl := index $.Db.ACL $uuid}}
{{- /* TODO -- --after-lb --meter */}}
ovn-nbctl acl-add --type=port-group {{if $acl.Log}}--log {{end}}{{if $acl.Severity}}--severity={{$acl.Severity}} {{end}}{{if $acl.Name}}--name={{$acl.Name}} {{end}}{{if $acl.Label}}--label={{$acl.Label}} {{end}}{{$pgName}} {{$acl.Direction}} {{$acl.Priority}} '{{$acl.Match}}' {{$acl.Action}}
{{- end }}
{{ end -}}