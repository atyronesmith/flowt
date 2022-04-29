
package {{.PkgName}}

// Type: {{.Schema.Type }}
// Version: {{.Schema.Version}}
// Tables: {{len .Schema.Tables}}

{{range .DBDef.Tables}}
// {{.JsonName}}
type {{.Name}} struct {
{{- range .Columns}}
    {{.Name}} {{.Type}} `json:"{{.JsonName}}"` {{ if .Comment}}// {{.Comment}} {{end}}
{{- end}}
}
{{end}}

type {{.DBDef.Name}} struct {
    Date Time `json:"_date"`
    Comment string `json:"_comment"`
    IsDiff bool `json:"_is_diff"`
    {{- range .DBDef.Tables}}
        {{.Name}} map[string]{{.Name}} `json:"{{.JsonName}}"`
    {{- end}}
}

{{ if eq .DBDef.Name "OVNSouthbound" }}
func (nb *OVNSouthbound) IsValid() bool {
	return len(nb.SBLogicalFlow) > 0 
}
{{else}}
func (nb *{{.DBDef.Name}}) IsValid() bool {
	return len(nb.NBLogicalSwitchPort) > 0 
}
{{end}}