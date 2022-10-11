package reflect

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/entity/api"
)

type fileDescriptor struct {
	*desc.FileDescriptor
}

// NewFileDescriptor
//
//	@param desc
//	@param rpc
//	@return fileDescriptor
func NewFileDescriptor(desc *desc.FileDescriptor) *fileDescriptor {
	return &fileDescriptor{
		FileDescriptor: desc,
	}
}

// GetAllApis
//
//	@receiver f
//	@return []api.ExportApis
//	@return error
func (f *fileDescriptor) GetAllApis() []api.ExportApis {
	svr := NewSvrReflect(f.GetServices())
	return svr.GetAipsInfo()
}
