package core

func FileContentMerge(content FileContent, merContent ...FileContent) FileContent {
	for _, item := range merContent {
		for key, value := range item {
			content[key] = value
		}
	}

	return content
}

// 把获取的配置项，转换成字符串切片
func GetListByCoreResult(key string, data interface{}) []string {
	list := data.(FileContent)[key].([]interface{})
	strArr := make([]string, len(list))

	for key, item := range list {
		strArr[key] = item.(string)
	}

	return strArr
}