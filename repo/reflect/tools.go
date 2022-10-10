package reflect

import "github.com/kalifun/proto-doc/entity/api"

// mapToComplexType
//
//	@param data
//	@return []api.ComplexType
func mapToComplexType(data map[string]api.ComplexType) []api.ComplexType {
	var c []api.ComplexType
	for _, v := range data {
		c = append(c, v)
	}
	return c
}

// mapToEnumType
//
//	@param data
//	@return []api.EnumType
func mapToEnumType(data map[string]api.EnumType) []api.EnumType {
	var e []api.EnumType
	for _, v := range data {
		e = append(e, v)
	}
	return e
}
