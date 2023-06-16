// Package conf
// @title
// @description
// @author njy
// @since 2022/11/21 13:17
package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

type tomlConfig struct {
	Mysql mysqlConfig `toml:"mysql"`
	Redis redisConfig `toml:"redis"`
	Agora agoraConfig `toml:"agora"`
	Amp   ampConfig   `toml:"amp"`
}
type mysqlConfig struct {
	Username string `toml:"username"`
	Pwd      string `toml:"pwd"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	DB       string `toml:"db"`
}
type redisConfig struct {
	Addr   string `toml:"addr"`
	DB     int    `toml:"db"`
	Secure bool   `toml:"secure"`
	Pwd    string `toml:"pwd"`
}
type agoraConfig struct {
	AppID          string `toml:"app_id"`
	AppCertificate string `toml:"app_certificate"`
}
type ampConfig struct {
	GeoKey string `toml:"geo_key"`
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	_, err := toml.DecodeFile("conf/conf.toml", &Cfg)
	if err != nil {
		log.Fatal("load system conf file failed:", err)
	}
}
