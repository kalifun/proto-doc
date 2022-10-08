package reflect

import (
	"fmt"

	"github.com/kalifun/proto-doc/entity/api"
	"gopkg.in/yaml.v3"
)

// ConversionRpcComments
//
//	@param msg
//	@return api.ApiDescribe
func ConversionRpcComments(msg string) api.ApiDescribe {
	var desc api.ApiDescribe
	comments := make(map[string]string)
	err := yaml.Unmarshal([]byte(msg), &comments)
	if err != nil {
		fmt.Println(err)
		return desc
	}
	if v, ok := comments["Path"]; ok {
		desc.Path = v
	}
	if v, ok := comments["Desc"]; ok {
		desc.Desc = v
	}
	if v, ok := comments["Version"]; ok {
		desc.Version = v
	}
	if v, ok := comments["Method"]; ok {
		desc.Method = v
	}
	return desc
}

// ConversionFieldComments
//
//	@param msg
//	@return api.Param
func ConversionFieldComments(msg string) api.Param {
	var param api.Param
	comments := make(map[string]string)
	err := yaml.Unmarshal([]byte(msg), &comments)
	if err != nil {
		param.Desc = msg
		return param
	}

	if v, ok := comments["Required"]; ok {
		param.Required = v
	} else {
		param.Required = "UNKNOW"
	}

	if v, ok := comments["Desc"]; ok {
		param.Desc = v
	}

	return param
}
