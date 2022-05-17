
package {{.PkgName}}

// Type: {{.Schema.Type }}
// Version: {{.Schema.Version}}
// Tables: {{len .Schema.Tables}}

{{range .DBDef.TableDefs}}
// {{.JsonName}}
type {{ postfix .Name $.Schema.Type.Postfix }} struct {
{{- range .Columns}}
    {{.Name}} {{.Type}} `json:"{{.JsonName}}{{ if .Optional}},omitempty{{end}}"` {{ if or .Comment .RefTable }}// {{.Comment}} {{.RefTable}} {{end}}
{{- end}}
}
{{end}}

type {{ .DBDef.Name }} struct {
    Date Time `json:"_date"`
    Comment string `json:"_comment"`
    IsDiff bool `json:"_is_diff"`
    {{- range .DBDef.TableDefs}}
        {{ .Name }} map[string]{{postfix .Name $.Schema.Type.Postfix }} `json:"{{.JsonName}}"`
    {{- end}}
}

{{ if eq .DBDef.Name "OVNSouthbound" }}
func (sb *OVNSouthbound) IsValid() bool {
	return len(sb.SBGlobal) > 0 
}
{{else}}
func (nb *{{.DBDef.Name}}) IsValid() bool {
	return len(nb.NBGlobal) > 0 
}
{{end}}