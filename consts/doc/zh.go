package doc

const MdZh = `
# {{.Name}}

### 接口描述

- 接口路径: {{.ApiDescribe.Path}}
- 接口描述: {{.ApiDescribe.Desc}}
- 访问限制: {{.ApiDescribe.Limit}}
- 请求方式: {{.ApiDescribe.Method}}
- 接口版本: {{.ApiDescribe.Version}}

### 输入参数

| 参数名称 | 是否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |

{{- range .InputParams}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}

### 输出参数

| 参数名称 | 是否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |

{{- range .OutputParams}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}

### 枚举值

{{- if .EnumParams}}
{{- range .EnumParams}}

#### {{.Name}}

| 枚举值 | 枚举注解 |
| ------ | -------- |

{{- range .Params}}
| {{.Name}} | {{.Desc}} |
{{- end}}
{{- end}}
{{- else}}

{{- end}}

### 复杂类型

{{- range .ComplexParams}}

#### {{.Name}}

> {{.Comments}}

| 参数名称 | 是否必选 | 参数类型 | 参数注解 |
| -------- | -------- | -------- | -------- |

{{- range .Params}}
| {{.Name}} | {{.Required}} | {{.Type}} | {{.Desc}} |
{{- end}}
{{- end}}
`
