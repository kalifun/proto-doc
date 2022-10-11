package reflect

import (
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/entity/api"
)

// svrReflect  TODO
type svrReflect struct {
	svrs []*desc.ServiceDescriptor
}

// NewSvrReflect
//
//	@return *svrReflect
func NewSvrReflect(svrs []*desc.ServiceDescriptor) *svrReflect {
	return &svrReflect{
		svrs: svrs,
	}
}

// GetAipsInfo
//
//	@receiver s
//	@return []api.ExportApis
func (s *svrReflect) GetAipsInfo() []api.ExportApis {
	var apis []api.ExportApis
	for _, svr := range s.svrs {
		apis = append(apis, FindMethods(svr))
	}
	return apis
}

// FindMethods
//
//	@param svr
//	@return []api.ApiInfo
func FindMethods(svr *desc.ServiceDescriptor) api.ExportApis {
	var export api.ExportApis
	for _, method := range svr.GetMethods() {
		m := NewMethodReflect(method)
		api := m.GetExportApi()
		export.Apis = append(export.Apis, api)
	}
	export.SvrName = svr.GetName()
	comments := svr.GetSourceInfo().GetLeadingComments()
	export.Comments = strings.ReplaceAll(strings.TrimSpace(comments), "\n", "")
	return export
}
