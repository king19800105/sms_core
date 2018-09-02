package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
)

type LogInfo struct {
	CreatedTime time.Time
}

func LogSave() gin.HandlerFunc {
	// todo... 如果没有启动队列，则直接记录数据库。如果没有启动数据库，就panic
	log := LogInfo{}

	return func(c *gin.Context) {
		log.CreatedTime = time.Now()
		fmt.Println(log)

		// before request

		c.Next()

		//// after request
		//latency := time.Since(t)
		//log.Print(latency)
		//
		//// access the status we are sending
		//status := c.Writer.Status()
		//log.Println(status)
	}
}