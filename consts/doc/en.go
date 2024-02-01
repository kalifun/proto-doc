package doc

const MdEn = `
# {{.Name}}

### Api Describe

- Api Path: {{.ApiDescribe.Path}}
- Api Description: {{.ApiDescribe.Desc}}
- Access Limitations: {{.ApiDescribe.Limit}}
- Request Method: {{.ApiDescribe.Method}}
- Api Version: {{.ApiDescribe.Version}}

### Input

| Parameter Name | Required or not | Parameter Type | Parameter annotation |
| -------- | -------- | -------- | -------- |

{{- range .InputParams}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}

### Output

| Parameter Name | Required or not | Parameter Type | Parameter annotation |
| -------- | -------- | -------- | -------- |

{{- range .OutputParams}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}

### Enums

{{- if .EnumParams}}
{{- range .EnumParams}}

#### {{.Name}}

| Enumerated values | Enumeration Notes |
| ------ | -------- |

{{- range .Params}}
| {{.Name}} | {{.Desc}} |
{{- end}}
{{- end}}
{{- else}}

{{- end}}

### Complex Types

{{- range .ComplexParams}}

#### {{.Name}}

> {{.Comments}}

| Parameter Name | Required or not | Parameter Type | Parameter annotation |
| -------- | -------- | -------- | -------- |

{{- range .Params}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}
{{- end}}
`
