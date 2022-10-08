package reflect

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/kalifun/proto-doc/entity/api"
)

// fileDescriptor  TODO
type fileDescriptor struct {
	*desc.FileDescriptor
	messages []*desc.MessageDescriptor
}

// NewFileDescriptor
//  @param desc
//  @return *fileDescriptor
func NewFileDescriptor(desc *desc.FileDescriptor) *fileDescriptor {
	return &fileDescriptor{
		FileDescriptor: desc,
	}
}

// GetApis
//  @receiver f
//  @return []api.ExportApis
//  @return error
func (f *fileDescriptor) GetApis() ([]api.ExportApis, error) {
	msg := GetAllMessage(f.FileDescriptor)
	f.messages = msg
	var apis []api.ExportApis
	for _, svr := range f.GetServices() {
		svrRef := NewSvrReflect(svr)
		methods := svrRef.FindMethod()
		var apiInfo api.ExportApis
		apiInfo.SvrName = svrRef.GetName()
		apiInfo.Comments = svrRef.GetComments()
		for _, method := range methods {
			var ai api.ApiInfo
			methodCom := ConversionRpcComments(method.comments)
			ai.ApiDescribe = methodCom
		}
	}
	return apis, nil
}

// GetAllMessage
//  @param f
//  @return messages
func GetAllMessage(f *desc.FileDescriptor) (messages []*desc.MessageDescriptor) {
	messages = append(messages, f.GetMessageTypes()...)

	for _, desc := range f.GetDependencies() {
		messages = append(messages, desc.GetMessageTypes()...)
	}
	return
}
