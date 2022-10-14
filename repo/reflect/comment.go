package reflect

import (
	"strings"

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
	text := TextPreProcessing(msg)
	err := yaml.Unmarshal([]byte(text), &comments)
	if err != nil {
		return desc
	}
	if v, ok := comments["path"]; ok {
		desc.Path = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
	}
	if v, ok := comments["desc"]; ok {
		desc.Desc = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
	}
	if v, ok := comments["version"]; ok {
		desc.Version = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
	}
	if v, ok := comments["method"]; ok {
		desc.Method = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
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
	text := TextPreProcessing(msg)
	err := yaml.Unmarshal([]byte(text), &comments)
	if err != nil {
		param.Desc = strings.ReplaceAll(strings.TrimSpace(msg), "\n", "")
		param.Required = "UNKNOW"
		return param
	}

	if v, ok := comments["required"]; ok {
		param.Required = v
	} else {
		param.Required = "UNKNOW"
	}

	if v, ok := comments["desc"]; ok {
		param.Desc = strings.ReplaceAll(strings.TrimSpace(v), "\n", "")
	}

	return param
}

// TextPreProcessing
//
//	@param msg
//	@return string
func TextPreProcessing(msg string) string {
	strs := []string{}
	for _, v := range strings.Split(msg, "\n") {
		strs = append(strs, strings.ToLower(strings.TrimPrefix(strings.TrimSpace(v), "@")))
	}

	return strings.Join(strs, "\n")
}
