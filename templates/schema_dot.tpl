digraph {
    concentrate=True;
    node [shape="none" fontsize="14"]
    edge [fontname="Helvetica,Arial,sans-serif"]

    LEGEND [tooltip="" label=<      
    <TABLE BORDER="2" COLOR="cornsilk2" CELLBORDER="0" CELLSPACING="0" CELLPADDING="3">
      <TR>
        <TD ALIGN="center" BGCOLOR="#E4E4E4" COLSPAN="2"><B>LEGEND</B></TD>
      </TR>
      <TR>
        <TD ALIGN="left">Optional</TD><TD>O</TD>
      </TR>
      <TR>
        <TD ALIGN="left">Index</TD><TD>I</TD>
      </TR>
      <TR>
        <TD ALIGN="left">Ephemeral</TD><TD>E</TD>
      </TR>
    </TABLE>
    >];

{{- range .DbDef.TableDefs}}
{{- $index := 0 }}
    {{.JsonName}} [tooltip="{{.ToolTip}}" label=<      
    <TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0" CELLPADDING="3">
      <TR>
        <TD ALIGN="center" PORT="input" BGCOLOR="#E4E4E4" COLSPAN="2"><B>{{.JsonName}}</B></TD>
      </TR>
      {{- range .Columns}}
      <TR>
        <TD ALIGN="left">{{.Name}}</TD><TD{{if .RefTable}} PORT="O{{$index}}"{{end}}>{{if .Index}}I{{end}}{{if .Optional}}O{{end}}{{if .Ephemeral}}E{{end}}{{if or .Index .Optional .Ephemeral}}{{else}}_{{end}}</TD>
      </TR>
{{- $index = add $index 1 -}} 
{{ end }}
    </TABLE>
    >];
{{- end}}

{{range .DbDef.TableDefs -}}{{$index := 0}}{{$tableName := .JsonName}}
{{- range .Columns -}}
    {{if .RefTable}}    {{$tableName}}:O{{$index}}:e -> {{.RefTable}}:input:c {{printf "\n"}}{{end}}
    {{- $index = add $index 1}}
{{- end -}}
{{end -}}
}
