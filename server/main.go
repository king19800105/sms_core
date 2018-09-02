package main

import (
	_ "gitee.com/king19800105/husms_sms_core/bootstrap"
	"gitee.com/king19800105/husms_sms_core/bootstrap"
	"github.com/gin-gonic/gin"
	"gitee.com/king19800105/husms_sms_core/routes"
)

func main() {
	r := gin.New()
	// 应用加载项
	bootstrap.Run(r)
	// 路由配置项
	routes.SetupApiRouter(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
