{
    "nodes": [
{{- $index := 0 -}}
{{range $i, $ele := .LogicalSwitch}}
  {{- if gt $index 0}},{{end}} {
    "itemStyle": {"color": "#4f19c7"},
    "name": "{{$ele.Name}}",
    "symbolSize": 50.0
    }{{- $index = add $index 1 -}}
{{- end -}}
{{range $i, $ele := .LogicalSwitchPort}}
  {{- if gt $index 0}},{{end}} {
    "itemStyle": {"color": "#19c7b9"},
    "name": "{{$i}}",
    "symbolSize": 4.5
    }{{- $index = add $index 1 -}}
{{- end -}}
    ],
    "links": [
{{- $index := 0 -}}
{{- range $i, $ele := .LogicalSwitch}}
{{- range $j, $ele2 := .Ports}}
  {{- if gt $index 0}},{{end}} {
    "source": "{{$ele.Name}}",
    "target": "{{$ele2}}"
    }{{- $index = add $index 1 -}}
{{- end -}}
{{- end -}}
    ]
}
    