// Package dao
// @title
// @description
// @author njy
// @since 2023/6/12 16:59
package dao

import "time"

type RtcServer struct {
	Id   int    `gorm:"column:id" db:"id" json:"id" form:"id"`
	Code string `gorm:"column:code" db:"code" json:"code" form:"code"` //服务器code，服务器厂商-地点-时间戳

	Core      int    `gorm:"column:core" db:"core" json:"core" form:"core"`                     //服务器cpu核心数
	Ram       int    `gorm:"column:ram" db:"ram" json:"ram" form:"ram"`                         //服务器内存大小（GB)
	Bandwidth int    `gorm:"column:bandwidth" db:"bandwidth" json:"bandwidth" form:"bandwidth"` //公网带宽（MB）
	Address   string `gorm:"column:address" db:"address" json:"address" form:"address"`         //服务器地址
	Location  string `gorm:"column:location" db:"location" json:"location" form:"location"`     //经纬度:lng,lat
	Operators int8   `gorm:"column:operators" db:"operators" json:"operators" form:"operators"` //服务器网络运营商
	Ipv4      string `gorm:"column:ipv4" db:"ipv4" json:"ipv4" form:"ipv4"`                     //公网IP

	AddTime    time.Time `gorm:"column:add_time" db:"add_time" json:"add_time" form:"add_time"`         //添加时间
	StartTime  time.Time `gorm:"column:start_time" db:"start_time" json:"start_time" form:"start_time"` //启用时间
	EndTime    time.Time `gorm:"column:end_time" db:"end_time" json:"end_time" form:"end_time"`         //过期时间
	Status     int8      `gorm:"column:status" db:"status" json:"status" form:"status"`                 //状态：1:配置中，2:启用，3:暂停，4:过期
	RtcType    int8      `gorm:"column:rtc_type" db:"rtc_type" json:"rtc_type" form:"rtc_type"`         //服务器类型：1:srs，
	NodeType   int8      `gorm:"column:node_type" db:"node_type" json:"node_type" form:"node_type"`     //节点类型：1:公网节点，2:内网节点
	Token      string    `gorm:"column:token" db:"token" json:"token" form:"token"`                     //节点token
	ConfigId   int       `gorm:"column:config_id" db:"config_id" json:"config_id" form:"config_id"`     //节点配置信息表id
	DomainId   int       `gorm:"column:domain_id" db:"domain_id" json:"domain_id" form:"domain_id"`     //节点域名信息配置表id
	UpdateTime time.Time `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
	Deleted    int8      `gorm:"column:deleted" db:"deleted" json:"deleted" form:"deleted"`
}
