package core

// 文件配置解析
type Parser interface {
	GetContentByCutPoint(string) (interface{}, bool)
	ReadFileByPath(string) FileContent
	SetFileContent(content FileContent) Parser
}

// 数据库实例链接
type DBConnection interface {
	initDBConnection(FileContent) DBConnection
}


// 缓存实例链接
type CacheConnection interface {
	initCacheConnection(FileContent) CacheConnection
}

// 队列实例
type QueueConnection interface {
	SetConfig(content FileContent) QueueConnection
	initProducer() QueueConnection
	InitConsumer(string, string) QueueConnection
	AddHandler(f interface{}) QueueConnection
}
