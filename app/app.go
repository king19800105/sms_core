package app

import (
	"reflect"
)

var configValue interface{}

// 参数1：容器中药获取的对象名称。参数2：对象转换后的类型
func App(name string, instanceType reflect.Type) {
	// 反射对象
}

//func GetConfig(path string) interface{} {
//	return getConfigByType("interface", path)
//}
//
//func GetConfigForString(path string) string {
//	return getConfigByType("string", path).(string)
//}
//
//func GetConfigForBool(path string) bool {
//	return getConfigByType("bool", path).(bool)
//}
//
//func GetConfigForInt(path string) int {
//	return getConfigByType("int", path).(int)
//}
//
//func GetConfigForList(path string) []interface{} {
//	return getConfigByType("slice", path).([]interface{})
//}
//
//func getConfigByType(parserType, path string) interface{} {
//	value, _ := core.GetConfig().GetContentByCutPoint(path)
//	ok := false
//
//	switch parserType {
//	case "string":
//		configValue, ok = value.(string)
//	case "bool":
//		configValue, ok = value.(bool)
//	case "int":
//		configValue, ok = value.(int)
//	case "slice":
//		configValue, ok = value.([]interface{})
//	case "interface":
//		configValue, ok = value.(interface{})
//	default:
//		configValue = nil
//	}
//
//	if nil == configValue || false == ok {
//		panic(message.ILLEGAL_TYPE)
//	}
//
//	return configValue
//}

