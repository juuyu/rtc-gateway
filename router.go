// Package main
// @title
// @description
// @author njy
// @since 2023/5/29 15:08
package main

import (
	"github.com/gin-gonic/gin"
	"rtc-gateway/agora"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	// agora服务
	agoraRouter := router.Group("agora")
	{
		agoraRouter.POST("/fetch_rtc_token", agora.RtcTokenHandler)
	}
	return router
}
