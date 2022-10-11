package reflect

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/entity/api"
)

type methodReflect struct {
	*desc.MethodDescriptor
}

// NewMethodReflect
//
//	@param method
//	@return *methodReflect
func NewMethodReflect(method *desc.MethodDescriptor) *methodReflect {
	return &methodReflect{
		MethodDescriptor: method,
	}
}

// GetExportApi
//
//	@receiver m
//	@return api.ApiInfo
func (m *methodReflect) GetExportApi() api.ApiInfo {
	var api api.ApiInfo
	ip := m.GetInputParams()
	api.InputParams = ip.Msg
	op := m.GetOutputParams()
	api.OutputParams = op.Msg
	api.ApiDescribe = m.GetApiDescribe()
	api.Name = m.GetName()
	api.ComplexParams = AggregateComplexTypes(ip.Dc.Complex, op.Dc.Complex)
	api.EnumParams = AggregateEnumTypes(ip.Dc.Enum, op.Dc.Enum)
	return api
}

// GetApiDescribe
//
//	@receiver m
//	@return api.ApiDescribe
func (m *methodReflect) GetApiDescribe() api.ApiDescribe {
	comments := m.GetSourceInfo().GetLeadingComments()
	return ConversionRpcComments(comments)
}

// GetInputParams
//
//	@receiver m
//	@return MsgDependences
func (m *methodReflect) GetInputParams() MsgDependences {
	input := m.GetInputType()
	msg := NewMsgReflect(input)
	return msg.GetMsgDependences()
}

// GetOutputParams
//
//	@receiver m
//	@return MsgDependences
func (m *methodReflect) GetOutputParams() MsgDependences {
	output := m.GetOutputType()
	msg := NewMsgReflect(output)
	return msg.GetMsgDependences()
}

// AggregateComplexTypes
//
//	@param types
//	@return []api.ComplexType
func AggregateComplexTypes(types ...map[string]api.ComplexType) []api.ComplexType {
	allC := make(map[string]api.ComplexType)
	ct := []api.ComplexType{}
	for _, c := range types {
		for ss, sc := range c {
			if _, ok := allC[ss]; !ok {
				allC[ss] = sc
				ct = append(ct, sc)
			}
		}
	}
	return ct
}

// AggregateEnumTypes
//
//	@param types
//	@return []api.EnumType
func AggregateEnumTypes(types ...map[string]api.EnumType) []api.EnumType {
	alle := make(map[string]api.EnumType)
	e := []api.EnumType{}
	for _, et := range types {
		for ss, se := range et {
			if _, ok := alle[ss]; !ok {
				alle[ss] = se
				e = append(e, se)
			}
		}
	}
	return e
}
