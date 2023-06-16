// Package api
// @title
// @description
// @author njy
// @since 2023/6/12 16:55
package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type AddServerParam struct {
	Core      int       `gorm:"column:core" db:"core" json:"core" form:"core"`                         //服务器cpu核心数
	Ram       int       `gorm:"column:ram" db:"ram" json:"ram" form:"ram"`                             //服务器内存大小（GB)
	Bandwidth int       `gorm:"column:bandwidth" db:"bandwidth" json:"bandwidth" form:"bandwidth"`     //公网带宽（MB）
	Address   string    `gorm:"column:address" db:"address" json:"address" form:"address"`             //服务器地址
	Operators int8      `gorm:"column:operators" db:"operators" json:"operators" form:"operators"`     //服务器网络运营商
	Ipv4      string    `gorm:"column:ipv4" db:"ipv4" json:"ipv4" form:"ipv4"`                         //公网IP
	StartTime time.Time `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"` //启用时间
	EndTime   time.Time `gorm:"column:end_time" db:"end_time" json:"end_time" form:"end_time"`         //过期时间
	Status    int8      `gorm:"column:status" db:"status" json:"status" form:"status"`                 //状态：1:配置中，2:启用，3:暂停，4:过期
	RtcType   int8      `gorm:"column:rtc_type" db:"rtc_type" json:"rtc_type" form:"rtc_type"`         //服务器类型：1:srs，
	NodeType  int8      `gorm:"column:node_type" db:"node_type" json:"node_type" form:"node_type"`     //节点类型：1:公网节点，2:内网节点
	ConfigId  int       `gorm:"column:config_id" db:"config_id" json:"config_id" form:"config_id"`     //节点配置信息表id
	DomainId  int       `gorm:"column:domain_id" db:"domain_id" json:"domain_id" form:"domain_id"`     //节点域名信息配置表id
}

// AddServer 添加服务器
func AddServer(ctx *gin.Context) {
	//
	var addServerParam AddServerParam
	err := ctx.BindJSON(&addServerParam)
	if err != nil {
		log.Println("bind json failed,err:", err)
		ctx.JSON(http.StatusBadRequest, "bind json failed")
	}
	//

	ctx.JSON(http.StatusOK, "ok")
}

// UpdateServer 修改服务器信息
func UpdateServer(ctx *gin.Context) {
}

// DelServer 删除服务器信息
func DelServer(ctx *gin.Context) {
}

// ListServer 获取全部服务器信息
func ListServer(ctx *gin.Context) {
}

// GetServerInfo 获取单个服务信息
func GetServerInfo(ctx *gin.Context) {
	//code := ctx.Param("code")

}
