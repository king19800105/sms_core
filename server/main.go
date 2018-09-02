package main

import (
	_ "github.com/king19800105/sms_core/bootstrap"
	"github.com/king19800105/sms_core/bootstrap"
	"github.com/gin-gonic/gin"
	"github.com/king19800105/sms_core/routes"
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
