package consts

type ParamType string

const (
	BOOL    ParamType = "Bool"
	BYTES   ParamType = "[]Byte"
	UINT32  ParamType = "Uint32"
	UINT64  ParamType = "Uint64"
	FLOAT   ParamType = "Float32"
	DOUBLE  ParamType = "Float64"
	INT32   ParamType = "Int32"
	INT64   ParamType = "INT64"
	STRING  ParamType = "String"
	ENUM    ParamType = "Enum"
	MESSAGE ParamType = "Msg"
	Group   ParamType = "Group"
	Null    ParamType = "Null"
)
