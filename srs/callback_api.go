// Package srs
// @title
// @description
// @author njy
// @since 2023/6/1 13:23
package srs

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @link https://ossrs.net/lts/zh-cn/docs/v4/doc/http-callback

type PublishData struct {
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	IP       string `json:"ip"`
	Vhost    string `json:"vhost"`
	App      string `json:"app"`
	Stream   string `json:"stream"`
	Param    string `json:"param"`
}

func Publish(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 0)
}

type UnpublishedData struct {
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	IP       string `json:"ip"`
	Vhost    string `json:"vhost"`
	App      string `json:"app"`
	Stream   string `json:"stream"`
	Param    string `json:"param"`
	PageURL  string `json:"pageUrl"`
}

func Unpublished(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 0)
}

type PlayData struct {
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	IP       string `json:"ip"`
	Vhost    string `json:"vhost"`
	App      string `json:"app"`
	Stream   string `json:"stream"`
	Param    string `json:"param"`
}

func Play(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 0)
}

type StopData struct {
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	IP       string `json:"ip"`
	Vhost    string `json:"vhost"`
	App      string `json:"app"`
	Stream   string `json:"stream"`
	Param    string `json:"param"`
	Cwd      string `json:"cwd"`
	File     string `json:"file"`
}

func Stop(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, 0)
}

type HlsData struct {
	ServerID string  `json:"server_id"`
	Action   string  `json:"action"`
	ClientID string  `json:"client_id"`
	IP       string  `json:"ip"`
	Vhost    string  `json:"vhost"`
	App      string  `json:"app"`
	Stream   string  `json:"stream"`
	Param    string  `json:"param"`
	Duration float64 `json:"duration"`
	Cwd      string  `json:"cwd"`
	File     string  `json:"file"`
	URL      string  `json:"url"`
	M3U8     string  `json:"m3u8"`
	M3U8URL  string  `json:"m3u8_url"`
	SeqNo    int     `json:"seq_no"`
}

func Hls(ctx *gin.Context) {
	var hlsParam HlsData
	err := ctx.BindJSON(&hlsParam)
	if err != nil {
		log.Println("bind json failed,err:", err)
	}
	log.Println("收到hls录制回调, param =", hlsParam)
	//go HlsSave(&hlsParam)
	ctx.JSON(http.StatusOK, 0)
}
