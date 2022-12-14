package api

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
