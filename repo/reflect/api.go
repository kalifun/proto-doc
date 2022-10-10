package reflect

import (
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/consts"
	"github.com/kalifun/proto-doc/entity/api"
)

// fileDescriptor  TODO
type fileDescriptor struct {
	*desc.FileDescriptor
	dc      *api.DataCollection
	complex []api.ComplexType
	enum    []api.EnumType
	boot    bool
}

// NewFileDescriptor
//  @param desc
//  @return *fileDescriptor
func NewFileDescriptor(desc *desc.FileDescriptor) *fileDescriptor {
	return &fileDescriptor{
		FileDescriptor: desc,
		dc: &api.DataCollection{
			Complex: make(map[string]api.ComplexType),
			Enum:    make(map[string]api.EnumType),
		},
	}
}

// GetAllApis
//  @receiver f
//  @return []api.ExportApis
//  @return error
func (f *fileDescriptor) GetAllApis() ([]api.ExportApis, error) {
	var apis []api.ExportApis
	for _, svr := range f.GetServices() {
		svrRef := NewSvrReflect(svr)
		methods := svrRef.FindMethod()
		var apiInfo api.ExportApis
		apiInfo.SvrName = svrRef.GetName()
		apiInfo.Comments = svrRef.GetComments()
		for _, method := range methods {
			var ai api.ApiInfo
			ai.Name = method.name
			methodCom := ConversionRpcComments(method.comments)
			ai.ApiDescribe = methodCom
			input := f.getParam(method.inputFullName)
			output := f.getParam(method.outputFullName)
			ai.InputParams = input
			ai.OutputParams = output
			// dc := f.GetComplex()
			dc := f.GetDataCollection()
			ai.ComplexParams = mapToComplexType(dc.Complex)
			ai.EnumParams = mapToEnumType(dc.Enum)
			apiInfo.Apis = append(apiInfo.Apis, ai)
		}
		apis = append(apis, apiInfo)
	}
	return apis, nil
}

// getParam
//  @receiver f
//  @param fullName
//  @return []api.Param
func (f *fileDescriptor) getParam(fullName string) []api.Param {
	msg := f.FindMessage(fullName)
	var params []api.Param
	for _, field := range msg.GetFields() {
		var typeName string
		fieldType := switchType(field.GetType())
		typeName = string(fieldType)

		if fieldType == consts.ENUM {
			typeName = field.GetName()
			enum := field.GetEnumType()
			if enum != nil {
				f.enum = append(f.enum, api.EnumType{
					Name: enum.GetFullyQualifiedName(),
				})
			}
		}

		if fieldType == consts.MESSAGE {
			typeName = field.GetName()
			msg := field.GetMessageType()
			if msg != nil {
				f.complex = append(f.complex, api.ComplexType{
					Name: msg.GetFullyQualifiedName(),
				})
			}
		}
		fieldCom := field.GetSourceInfo().GetLeadingComments()
		com := ConversionFieldComments(fieldCom)
		param := api.Param{
			Name:     field.GetName(),
			Required: com.Required,
			Type:     GetParamType(typeName, field.GetLabel()),
			Desc:     com.Desc,
		}
		params = append(params, param)
	}
	return params
}

// GetDataCollection
//  @receiver f
//  @return *api.DataCollection
func (f *fileDescriptor) GetDataCollection() *api.DataCollection {
	if f.boot {
		return f.dc
	}
	for _, v := range f.complex {
		msg := f.findMessage(v.Name)
		if msg != nil {
			f.startDataCollection(msg)
		}
	}
	for _, v := range f.enum {
		enum := f.findEnum(v.Name)
		if enum != nil {
			if _, ok := f.dc.Enum[enum.GetName()]; !ok {
				f.dc.Enum[enum.GetName()] = f.collectionEnum(enum)
			}
		}
	}
	f.boot = true
	return f.dc
}

// findMessage
//  @receiver f
//  @param fullName
//  @return *desc.MessageDescriptor
func (f *fileDescriptor) findMessage(fullName string) *desc.MessageDescriptor {
	msg := f.FindMessage(fullName)
	if msg != nil {
		return msg
	}
	for _, v := range f.GetDependencies() {
		msg := v.FindMessage(fullName)
		if msg != nil {
			return msg
		}
	}
	return nil
}

// startDataCollection
//  @receiver f
//  @param msg
func (f *fileDescriptor) startDataCollection(msg *desc.MessageDescriptor) {
	var c api.ComplexType
	c.Name = msg.GetName()
	c.Comments = strings.ReplaceAll(strings.TrimSpace(msg.GetSourceInfo().GetLeadingComments()), "\n", "")
	for _, field := range msg.GetFields() {
		var typeName string
		fieldType := switchType(field.GetType())
		typeName = string(fieldType)
		if fieldType == consts.ENUM {
			enum := field.GetEnumType()
			if enum != nil {
				typeName = msg.GetName()
				if _, ok := f.dc.Enum[enum.GetName()]; !ok {
					f.dc.Enum[enum.GetName()] = f.collectionEnum(enum)
				}
			}
		}

		if fieldType == consts.MESSAGE {
			msg := field.GetMessageType()
			if msg != nil {
				typeName = msg.GetName()
				f.startDataCollection(msg)
			}
		}
		fieldCom := field.GetSourceInfo().GetLeadingComments()
		com := ConversionFieldComments(fieldCom)
		param := api.Param{
			Name:     field.GetName(),
			Required: com.Required,
			Type:     GetParamType(typeName, field.GetLabel()),
			Desc:     com.Desc,
		}
		c.Params = append(c.Params, param)
	}

	if _, ok := f.dc.Complex[c.Name]; !ok {
		f.dc.Complex[c.Name] = c
	}
}

// collectionEnum 
//  @receiver f 
//  @param enum 
//  @return api.EnumType 
func (f *fileDescriptor) collectionEnum(enum *desc.EnumDescriptor) api.EnumType {
	var et api.EnumType
	et.Name = enum.GetName()
	for _, e := range enum.GetValues() {
		var ep api.EnumParam
		ep.Name = e.GetName()
		ep.Desc = strings.ReplaceAll(strings.TrimSpace(e.GetSourceInfo().GetLeadingComments()), "\n", "")
		et.Params = append(et.Params, ep)
	}
	return et
}

// findEnum
//  @receiver f
//  @param fullName
//  @return *desc.EnumDescriptor
func (f *fileDescriptor) findEnum(fullName string) *desc.EnumDescriptor {
	enum := f.FindEnum(fullName)
	if enum != nil {
		return enum
	}
	for _, v := range f.GetDependencies() {
		enum := v.FindEnum(fullName)
		if enum != nil {
			return enum
		}
	}
	return nil
}
