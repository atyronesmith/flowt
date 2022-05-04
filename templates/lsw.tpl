{
    "nodes": [
{{- $index := 0 -}}
{{range $i, $ele := .LogicalSwitch}}
  {{- if gt $index 0}},{{end}} {
    "itemStyle": {"color": "#4f19c7"},
    "name": "{{$ele.Name}}",
    "symbolSize": {{len $ele.Ports}}
    }{{- $index = add $index 1 -}}
{{- end -}}
{{range $i, $ele := .LogicalSwitchPort}}
  {{- if gt $index 0}},{{end}} {
    "itemStyle": {"color": "#4f19c7"},
    "name": "{{$ele.Name}}",
    "symbolSize": 1
    }{{- $index = add $index 1 -}}
{{- end -}}
    ]
}
    