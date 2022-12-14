package proto

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/entity/api"
	"github.com/kalifun/proto-doc/repo/reflect"
)

// ReflectProtos
//
//	@param files
//	@return []api.ExportApis
//	@return error
func ReflectProtos(files ...*desc.FileDescriptor) []api.ExportApis {
	var apis []api.ExportApis
	for _, file := range files {
		desc := reflect.NewFileDescriptor(file)
		descApi := desc.GetAllApis()
		apis = append(apis, descApi...)
	}
	return apis
}
