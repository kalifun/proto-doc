package reflect

import (
	"fmt"

	"github.com/kalifun/proto-doc/consts"
	"google.golang.org/protobuf/types/descriptorpb"
)

// switchType
//
//	@param ft
//	@return consts.ParamType
func switchType(ft descriptorpb.FieldDescriptorProto_Type) consts.ParamType {
	var typeName consts.ParamType
	switch ft {
	case descriptorpb.FieldDescriptorProto_TYPE_BOOL:
		typeName = consts.BOOL
	case descriptorpb.FieldDescriptorProto_TYPE_BYTES:
		typeName = consts.BYTES
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED32:
		typeName = consts.UINT32
	case descriptorpb.FieldDescriptorProto_TYPE_FIXED64:
		typeName = consts.UINT64
	case descriptorpb.FieldDescriptorProto_TYPE_FLOAT:
		typeName = consts.FLOAT
	case descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:
		typeName = consts.DOUBLE
	case descriptorpb.FieldDescriptorProto_TYPE_SINT32:
		typeName = consts.INT32
	case descriptorpb.FieldDescriptorProto_TYPE_SINT64:
		typeName = consts.INT64
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED32:
		typeName = consts.INT32
	case descriptorpb.FieldDescriptorProto_TYPE_SFIXED64:
		typeName = consts.INT64
	case descriptorpb.FieldDescriptorProto_TYPE_INT32:
		typeName = consts.INT32
	case descriptorpb.FieldDescriptorProto_TYPE_INT64:
		typeName = consts.INT64
	case descriptorpb.FieldDescriptorProto_TYPE_STRING:
		typeName = consts.STRING
	case descriptorpb.FieldDescriptorProto_TYPE_UINT32:
		typeName = consts.UINT32
	case descriptorpb.FieldDescriptorProto_TYPE_UINT64:
		typeName = consts.UINT64
	case descriptorpb.FieldDescriptorProto_TYPE_ENUM:
		typeName = consts.ENUM
	case descriptorpb.FieldDescriptorProto_TYPE_MESSAGE:
		typeName = consts.MESSAGE
	case descriptorpb.FieldDescriptorProto_TYPE_GROUP:
		typeName = consts.Group
	default:
		typeName = consts.Null
	}

	return typeName
}

// GetParamType
//
//	@param paramType
//	@param label
//	@return string
func GetParamType(paramType consts.ParamType, label descriptorpb.FieldDescriptorProto_Label) string {
	var typeName string
	if paramType != consts.Null {
		if label == descriptorpb.FieldDescriptorProto_LABEL_REPEATED {
			typeName = fmt.Sprintf("Array of %s", string(typeName))
		}
	}
	return string(typeName)
}
