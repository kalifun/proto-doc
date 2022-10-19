package export

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/kalifun/proto-doc/entity/api"
)

// markdown
type markdown struct {
	template string // 配置新模板
	outPath  string
	language string
}

type option func(*markdown)

// SetTemplate 设置模板
//  @param template
//  @return option
func SetTemplate(template string) option {
	return func(m *markdown) {
		m.template = template
	}
}

// SetOutPath 设置输出路径
//  @param outPath
//  @return option
func SetOutPath(outPath string) option {
	return func(m *markdown) {
		m.outPath = outPath
	}
}

// SetLanguage
//  @param lan
//  @return option
func SetLanguage(lan string) option {
	return func(m *markdown) {
		m.language = lan
	}
}

// NewMarkdown
//  @param opts
//  @return *markdown
func NewMarkdown(opts ...option) *markdown {
	m := markdown{}
	for _, f := range opts {
		f(&m)
	}
	return &m
}

// GenMarkdown 
//  @receiver m 
//  @param apis 
//  @return error 
func (m *markdown) GenMarkdown(apis []api.ExportApis) error {
	m.setTemplate()
	for _, api := range apis {
		for _, apiInfo := range api.Apis {
			fileName := fmt.Sprintf("%s.md", apiInfo.Name)
			f, err := os.Create(path.Join(m.getOutPath(), fileName))
			if err != nil {
				return err
			}
			tl, err := template.New("doc").Parse(m.template)
			if err != nil {
				return err
			}
			w := bufio.NewWriter(f)
			err = tl.Execute(w, apiInfo)
			if err != nil {
				return err
			}
			w.Flush()
		}
	}
	return nil
}

// setTemplate
//  @receiver m
func (m *markdown) setTemplate() {
	if m.template == "" {
		// 判断用什么模板返回
		switch m.language {
		case "zh":
			m.template = mdZh
		case "en":
			m.template = mdEn
		default:
			m.template = mdEn
		}
	}
}

// getOutPath
//  @receiver m
//  @return string
func (m *markdown) getOutPath() string {
	if m.outPath == "" {
		return "./"
	}
	return m.outPath
}

const mdZh = `
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

const mdEn = `
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
