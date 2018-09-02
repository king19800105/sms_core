package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/king19800105/sms_core/app/http/controllers"
)

// 发送api ：http://www.husms.com/sms/api
// 余量接口 ：http://www.husms.com/sms/query

func SetupApiRouter(r *gin.Engine) *gin.Engine {
	api := r.Group("/api")
	{
		api.POST("/sms-send", controllers.SMSSend)
		api.GET("/sms-query", controllers.SMSQuery)
	}

	return r
}