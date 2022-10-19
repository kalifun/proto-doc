# proto doc

You can generate `markdown` documents from `proto`'s comments

Comments have their own set of criteria, if you want to try it you can visit this [repository](https://github.com/kalifun/vscode-proto3-tools.git)

## how to use

```bash
❯ ./proto-doc -h
proto-doc is a tool for generating documentation and annotations from proto files

Usage:
  proto-doc [flags]
  proto-doc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         generating documentation
  help        Help about any command

Flags:
  -h, --help   help for proto-doc

Use "proto-doc [command] --help" for more information about a command.
```
Try to generate a document

```bash
❯ ./proto-doc doc -h
Parsing proto to generate markdown documents

Usage:
  proto-doc doc [flags]

Flags:
  -h, --help              help for doc
      --language string   export document template language (default is en) (default "en")
      --out string        export document path (default is .) (default ".")
      --proto string      proto files (default is .) (default ".")
```

```bash
./proto-doc doc --proto xxx.proto --out docs --language en
```

## feature

### markdown template

```template
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
```

## 💡ideas

This tool attempts to parse the proto file, which will eventually produce a relatively readable structure. If you are not satisfied with the default template, you can use a structure to implement other functions.

```go
// 导出的接口
type ExportApis struct {
	SvrName  string    // 服务名称
	Comments string    // 服务描述
	Apis     []ApiInfo // 此服务下的所有接口
}

// 接口导出
type ApiInfo struct {
	Name          string        // 接口名称
	ApiDescribe   ApiDescribe   // 接口描述
	InputParams   []Param       // 输入参数
	InputDemo     string        // 输出示例
	OutputParams  []Param       // 输出参数
	OutputDemo    string        // 输出示例
	ComplexParams []ComplexType // 复杂参数
	EnumParams    []EnumType    // 枚举值
	ErrorCode     []ErrorMsg    // 错误码
}

// 接口描述
type ApiDescribe struct {
	Path    string // 接口路径
	Desc    string // 接口描述
	Limit   string // 频率限制
	Source  string // 请求来源
	Method  string // 请求方式
	Version string // 版本
}

// 参数
type Param struct {
	Name     string // 参数名称
	Required string // 是否必选
	Type     string // 参数类型
	Desc     string // 参数描述
}

// 复杂类型
type ComplexType struct {
	Name     string  // 复杂类型名称
	Comments string  // 复杂类型描述
	Params   []Param // 复杂类型参数
}

// 错误码
type ErrorMsg struct {
	Code    string // 错误码
	CodeMsg string // 错误信息
	Desc    string // 错误描述
}

// 枚举类型
type EnumType struct {
	Name     string      // 枚举类型名称
	Comments string      // 枚举类型描述
	Params   []EnumParam // 枚举类型参数
}

// 枚举值
type EnumParam struct {
	Name string // 枚举值
	Desc string // 参数描述
}

// 数据收集器
type DataCollection struct {
	Complex map[string]ComplexType // 复杂类型
	Enum    map[string]EnumType    // 枚举类型
}

```