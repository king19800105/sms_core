package bootstrap

import (
	"gitee.com/king19800105/husms_sms_core/core"
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	// 加载核心模块
	core.Load()
	// 设置模式
	// setMode()
	// 设置日志生成方式
	// setLogSaveType(r)
}

//func setMode() {
//	mode := core.GetDebugMode()
//
//	if mode {
//		gin.SetMode(gin.DebugMode)
//		return
//	}
//
//	gin.SetMode(gin.ReleaseMode)
//}
//
//func setLogSaveType(r *gin.Engine) {
//	coreLog := core.GetLogSaveType()
//	gin.DisableConsoleColor()
//
//	switch coreLog.SaveType {
//	case "file":
//		setLogToFile(r, coreLog)
//	case "db":
//		setLogToDB(r)
//	default:
//		panic(message.ILLEGAL_LOG_TYPE)
//	}
//}
//
//func setLogToFile(route *gin.Engine, core core.LogInfo) {
//	req, reqOk := os.Create(core.RequestPath)
//	err, errOk := os.Create(core.ErrorPath)
//
//	if nil != reqOk && nil != errOk {
//		panic(message.LOG_FILE_CREATE_FAILED)
//	}
//
//	gin.DefaultWriter = io.MultiWriter(req)
//	gin.DefaultErrorWriter = io.MultiWriter(err)
//	route.Use(gin.Logger())
//	route.Use(gin.Recovery())
//}
//
//func setLogToDB(route *gin.Engine) {
//	gin.DefaultWriter = ioutil.Discard
//	gin.DefaultErrorWriter = ioutil.Discard
//	route.Use(middleware.LogSave())
//	route.Use(gin.Recovery())
//}
