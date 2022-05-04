
package {{.PkgName}}

// Type: {{.Schema.Type }}
// Version: {{.Schema.Version}}
// Tables: {{len .Schema.Tables}}

{{range .DBDef.Tables}}
// {{.JsonName}}
type {{ postfix .Name $.Schema.Type.Postfix }} struct {
{{- range .Columns}}
    {{.Name}} {{.Type}} `json:"{{.JsonName}}"` {{ if .Comment}}// {{.Comment}} {{end}}
{{- end}}
}
{{end}}

type {{ .DBDef.Name }} struct {
    Date Time `json:"_date"`
    Comment string `json:"_comment"`
    IsDiff bool `json:"_is_diff"`
    {{- range .DBDef.Tables}}
        {{ .Name }} map[string]{{postfix .Name $.Schema.Type.Postfix }} `json:"{{.JsonName}}"`
    {{- end}}
}

{{ if eq .DBDef.Name "OVNSouthbound" }}
func (nb *OVNSouthbound) IsValid() bool {
	return len(nb.LogicalFlow) > 0 
}
{{else}}
func (nb *{{.DBDef.Name}}) IsValid() bool {
	return len(nb.LogicalSwitchPort) > 0 
}
{{end}}