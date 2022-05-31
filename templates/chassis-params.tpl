controllers:
{{- range $chassisId := .Controllers -}}
{{$chassis := index $.SBDb.Chassis $chassisId}}
  - hostname: "{{$chassis.Hostname}}"
    ovn-bridge-mappings: "{{ index $chassis.ExternalIds "ovn-bridge-mappings"}}"
{{- end }}
computes:
{{- range $chassisId := .Computes -}}
{{$chassis := index $.SBDb.Chassis $chassisId}}
  - hostname: "{{$chassis.Hostname}}"
    ovn-bridge-mappings: "{{ index $chassis.ExternalIds "ovn-bridge-mappings"}}"
{{- end }}

{{$inactivity_probe := 180000}}
{{- range $key, $value := .SBDb.Connection -}}
{{ $inactivity_probe = $value.InactivityProbe}}
{{- end -}}
sb_inactivity_probe: {{if $inactivity_probe}}{{$inactivity_probe}}{{else}}180000{{end}}
{{$inactivity_probe := 180000}}
{{- range $key, $value := .Db.Connection -}}
{{ $inactivity_probe = $value.InactivityProbe}}
{{- end -}}
nb_inactivity_probe: {{if $inactivity_probe}}{{$inactivity_probe}}{{else}}180000{{end}}
