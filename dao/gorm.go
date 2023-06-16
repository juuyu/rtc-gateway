// Package dao
// @title
// @description
// @author njy
// @since 2023/6/13 14:01
package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"rtc-gateway/conf"
)

const (
	DsnFormat = "%v:%v@tcp(%v:%d)/%v?charset=utf8&parseTime=True&loc=Local"
)

var DB *gorm.DB

func InitMySQL() {
	//配置MySQL连接参数
	dsn := fmt.Sprintf(DsnFormat, conf.Cfg.Mysql.Username, conf.Cfg.Mysql.Pwd, conf.Cfg.Mysql.Host, conf.Cfg.Mysql.Port, conf.Cfg.Mysql.DB)
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败, err:", err)
	}
	DB = db
}
