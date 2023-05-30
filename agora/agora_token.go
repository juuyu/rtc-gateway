// Package agora
// @title
// @description
// @author njy
// @since 2023/5/30 11:57
package agora

import (
	rtcTokenBuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"rtc-gateway/conf"
)

const (
	// TokenExpireTimeInSeconds AccessToken2 过期的时间，单位为秒
	// 当 AccessToken2 过期但权限未过期时，用户仍在频道里并且可以发流，不会触发 SDK 回调。
	// 但一旦用户和频道断开连接，用户将无法使用该 Token 加入同一频道。请确保 AccessToken2 的过期时间晚于权限过期时间。
	TokenExpireTimeInSeconds = uint32(1000)
	// PrivilegeExpireTimeInSeconds 权限过期的时间，单位为秒。
	// 权限过期30秒前会触发 token-privilege-will-expire 回调。
	// 权限过期时会触发 token-privilege-did-expire 回调。
	// 为作演示，在此将过期时间设为 40 秒。你可以看到客户端自动更新 Token 的过程
	PrivilegeExpireTimeInSeconds = uint32(1000)
)

type rtcIntTokenStruct struct {
	UidRtcInt   uint32 `json:"uid"`
	ChannelName string `json:"ChannelName"`
	Role        uint32 `json:"role"`
}

func RtcTokenHandler(ctx *gin.Context) {
	var rtcOptions rtcIntTokenStruct
	var role rtcTokenBuilder.Role
	err := ctx.BindJSON(&rtcOptions)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "400")
	}
	uid := rtcOptions.UidRtcInt
	channel := rtcOptions.ChannelName

	switch rtcOptions.Role {
	case 1:
		role = rtcTokenBuilder.RolePublisher
	case 2:
		role = rtcTokenBuilder.RoleSubscriber
	}

	token := generateRtcToken(&uid, &channel, &role)
	ctx.JSON(http.StatusOK, *token)
}

func generateRtcToken(uid *uint32, channel *string, role *rtcTokenBuilder.Role) *string {
	result, err := rtcTokenBuilder.BuildTokenWithUid(conf.Cfg.Agora.AppID, conf.Cfg.Agora.AppCertificate, *channel, *uid, *role, TokenExpireTimeInSeconds, PrivilegeExpireTimeInSeconds)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Token with uid: ", result)
		log.Println("uid is: ", *uid)
		log.Println("ChannelName is: ", *channel)
		log.Println("Role is: ", *role)
	}
	return &result
}
