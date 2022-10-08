package reflect

import "github.com/jhump/protoreflect/desc"

// svrReflect  TODO 
type svrReflect struct {
	svr *desc.ServiceDescriptor
}

// Method  TODO 
type Method struct {
	name           string
	inputFullName  string
	outputFullName string
	comments       string
}

// NewSvrReflect
//
//	@param service
//	@return *svrReflect
func NewSvrReflect(service *desc.ServiceDescriptor) *svrReflect {
	return &svrReflect{
		svr: service,
	}
}

// FindMethod
//
//	@receiver s
func (s *svrReflect) FindMethod() []Method {
	var methods []Method
	for _, method := range s.svr.GetMethods() {
		name := method.GetName()
		input := method.GetInputType()
		output := method.GetOutputType()
		inputFullName := input.GetFullyQualifiedName()
		outputFullName := output.GetFullyQualifiedName()
		comments := method.GetSourceInfo().GetLeadingComments()
		methods = append(methods, Method{
			name:           name,
			inputFullName:  inputFullName,
			outputFullName: outputFullName,
			comments:       comments,
		})
	}
	return methods
}

// GetName
//
//	@receiver s
//	@return string
func (s *svrReflect) GetName() string {
	return s.svr.GetName()
}

// GetComments
//
//	@receiver s
//	@return string
func (s *svrReflect) GetComments() string {
	return s.svr.GetSourceInfo().GetLeadingComments()
}
