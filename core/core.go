package core

import (
	"github.com/sarulabs/di"
	"github.com/king19800105/sms_core/app/providers"
)

const CORE_FILE = "config/core.yml"

// 核心结构
type Core struct {
	container  di.Container
	db         DBConnection
	cache      CacheConnection
	queue      QueueConnection
	resource   FileContent
	fileParser Parser
}

var core *Core

// 映射
var constructMapping = map[string]interface{}{
	"parser": NewYmlParser,
	"mysql":  NewMysqlConnection,
	"redis":  NewRedisConnection,
	"nsq":    NewNSQConnection,
}

// 初始化核心
func init() {
	parser, ok := constructMapping["parser"]

	if !ok {
		panic("core parser load failed")
	}

	core = NewCore(parser)
	coreResult := core.fileParser.ReadFileByPath(CORE_FILE)

	if !ok {
		panic("core configuration file initialization failed")
	}

	core.resource["core"] = coreResult
}

// 基础数据加载
func Load() {
	core.loadAllYAMLInfo().
		loadDBConnection().
		loadCacheConnection().
		loadQueueConnection().
		loadContainer()
}


// Core对象构建
func NewCore(parser interface{}) *Core {
	return &Core{
		resource:   make(FileContent),
		fileParser: parser.(func() Parser)(),
	}
}

// 加载YAML配置文件
func (c *Core) loadAllYAMLInfo() *Core {
	pathList, ok := c.fileParser.GetContentByCutPoint("core.path")

	if !ok {
		panic("core path config parser failed")
	}

	c.resource["env"] = core.fileParser.ReadFileByPath(pathList.(FileContent)["env-file"].(string))
	c.resource["config"] = core.fileParser.ReadFileByPath(pathList.(FileContent)["config-dir"].(string))
	c.resource["i18n"] = core.fileParser.ReadFileByPath(pathList.(FileContent)["i18n-dir"].(string))

	return c
}

// 加载数据库链接对象
func (c *Core) loadDBConnection() *Core {
	allContent := c.getInstanceBaseInfo("db")
	db := constructMapping[allContent["driver"].(string)]
	c.db = db.(func() DBConnection)().
		initDBConnection(allContent)

	return c
}

// 加载缓存链接对象
func (c *Core) loadCacheConnection() *Core {
	allContent := c.getInstanceBaseInfo("cache")
	cache := constructMapping[allContent["driver"].(string)]
	c.cache = cache.(func() CacheConnection)().
		initCacheConnection(allContent)

	return c
}

// 加载队列
func (c *Core) loadQueueConnection() *Core {
	allContent := c.getInstanceBaseInfo("queue")
	queue := constructMapping[allContent["driver"].(string)]
	c.queue = queue.(func() QueueConnection)().
		SetConfig(allContent).
		initProducer()

	return c
}

// 容器初始化
func (c *Core) loadContainer() *Core {
	list := GetListByCoreResult("container-scope", c.resource["core"].(FileContent)["core"])
	containerBuild, err := di.NewBuilder(list...)

	if nil != err {
		panic("core container initialization failed")
	}

	containerBuild.Add(c.setContainer()...)
	c.container = containerBuild.Build()

	return c
}

// 整合容器定义
func (c *Core) setContainer() []di.Def {
	coreDef := []di.Def{
		{
			Name:  "resource",
			Scope: "core",
			Build: func(ctn di.Container) (interface{}, error) {
				return c.resource, nil
			},
		},
		{
			Name:  "db-connection",
			Scope: "core",
			Build: func(ctn di.Container) (interface{}, error) {
				return c.db, nil
			},
		},
		{
			Name:  "cache-connection",
			Scope: "core",
			Build: func(ctn di.Container) (interface{}, error) {
				return c.cache, nil
			},
		},
		{
			Name:  "queue-connection",
			Scope: "core",
			Build: func(ctn di.Container) (interface{}, error) {
				return c.queue, nil
			},
		},
		{
			Name:  "file-parser",
			Scope: "core",
			Build: func(ctn di.Container) (interface{}, error) {
				return c.fileParser, nil
			},
		},
	}

	app := providers.AppProvider{}.Register()

	if len(app) > 0 {
		coreDef = append(coreDef, app...)
	}

	return coreDef
}

// 设置链接参数列表
func (c *Core) getInstanceBaseInfo(instanceType string) FileContent {
	envContent, eOk := c.fileParser.SetFileContent(c.resource["env"].(FileContent)).GetContentByCutPoint(instanceType)
	coreContent, cOk := c.fileParser.SetFileContent(c.resource["core"].(FileContent)).GetContentByCutPoint("core." + instanceType)

	if !eOk || !cOk {
		panic("instance base info parser failed")
	}

	return FileContentMerge(envContent.(FileContent), coreContent.(FileContent))
}
