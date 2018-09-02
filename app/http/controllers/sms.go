package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 发送
func SMSSend(c *gin.Context) {
	c.String(http.StatusOK, "sms-send")
}

// 余额查询
func SMSQuery(c *gin.Context) {
	panic("222")
	c.String(http.StatusOK, "sms-query")
}