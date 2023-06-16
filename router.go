// Package main
// @title
// @description
// @author njy
// @since 2023/5/29 15:08
package main

import (
	"github.com/gin-gonic/gin"
	"rtc-gateway/agora"
	"rtc-gateway/srs"
	"rtc-gateway/web/api"
)

func StartRouter() *gin.Engine {
	router := gin.Default()
	// server
	serverRouter := router.Group("api/server/v1")
	{
		serverRouter.GET("/server", api.ListServer)
		serverRouter.GET("/server/:code", api.GetServerInfo)
	}
	{
		serverRouter.PUT("/server", api.UpdateServer)
	}
	{
		serverRouter.DELETE("/server/:code", api.DelServer)
	}
	{
		serverRouter.POST("/server", api.AddServer)
		serverRouter.POST("/fetch_forward_server")
		serverRouter.POST("/fetch_publish_server")
	}

	// srs
	srsRouter := router.Group("api/srs/v1")
	{
		// srs节点心跳
		srsRouter.POST("/ping", srs.Ping)
		// 获取配置文件
		srsRouter.POST("/fetch_conf", srs.FetchConf)
		// srs回调统一处理
		srsRouter.POST("/publish", srs.Publish)
		srsRouter.POST("/unpublished", srs.Unpublished)
		srsRouter.POST("/play", srs.Play)
		srsRouter.POST("/stop", srs.Stop)
		srsRouter.POST("/hls", srs.Hls)
	}

	// agora服务
	agoraRouter := router.Group("api/agora/v1")
	{
		agoraRouter.POST("/fetch_rtc_token", agora.RtcTokenHandler)
	}
	return router
}
