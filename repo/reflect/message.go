package reflect

import (
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/consts"
	"github.com/kalifun/proto-doc/entity/api"
)

type msgReflect struct {
	*desc.MessageDescriptor
	md   *MsgDependences
	boot bool
}

type collectionType int

const (
	current   collectionType = 1 // 获取本层次的
	recursion collectionType = 2 // 获取层次更深的
)

// NewMsgReflect
//
//	@param msg
//	@return *msgReflect
func NewMsgReflect(msg *desc.MessageDescriptor) *msgReflect {
	return &msgReflect{
		MessageDescriptor: msg,
		md: &MsgDependences{
			Msg: []api.Param{},
			Dc: api.DataCollection{
				Complex: make(map[string]api.ComplexType),
				Enum:    make(map[string]api.EnumType),
			},
		},
	}
}

// MsgDependences  信息依赖
type MsgDependences struct {
	Msg []api.Param        // 当前层的所有参数
	Dc  api.DataCollection // 信息依赖
}

// GetMsgDependences
//
//	@receiver m
//	@return MsgDependences
func (m *msgReflect) GetMsgDependences() MsgDependences {
	if m.boot {
		return *m.md
	}
	m.msgCollection(m.MessageDescriptor, current)
	return *m.md
}

// msgCollection
//
//	@receiver m
//	@param msg
//	@param ct
func (m *msgReflect) msgCollection(msg *desc.MessageDescriptor, ct collectionType) {
	var params []api.Param
	for _, field := range msg.GetFields() {
		var typeName string
		fieldType := switchType(field.GetType())
		typeName = string(fieldType)

		if fieldType == consts.ENUM {
			typeName = field.GetName()
			enum := field.GetEnumType()
			if enum != nil {
				if _, ok := m.md.Dc.Enum[enum.GetName()]; !ok {
					m.md.Dc.Enum[enum.GetName()] = m.enumCollection(enum)
				}
			}
		}

		if fieldType == consts.MESSAGE {
			typeName = field.GetName()
			msg := field.GetMessageType()
			if msg != nil {
				typeName = msg.GetName()
				m.msgCollection(msg, recursion)
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

	// 优先判别是属于哪种类型
	if ct == recursion {
		var c api.ComplexType
		c.Name = msg.GetName()
		c.Comments = strings.ReplaceAll(strings.TrimSpace(msg.GetSourceInfo().GetLeadingComments()), "\n", "")
		c.Params = params
		if _, ok := m.md.Dc.Complex[c.Name]; !ok {
			m.md.Dc.Complex[c.Name] = c
		}
	} else {
		m.md.Msg = append(m.md.Msg, params...)
	}
}

// enumCollection
//
//	@receiver m
//	@param enum
//	@return api.EnumType
func (m *msgReflect) enumCollection(enum *desc.EnumDescriptor) api.EnumType {
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
