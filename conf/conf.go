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
	Agora agoraConfig `toml:"agora"`
}

type agoraConfig struct {
	AppID          string `toml:"app_id"`
	AppCertificate string `toml:"app_certificate"`
}

var Cfg *tomlConfig

func init() {
	Cfg = new(tomlConfig)
	_, err := toml.DecodeFile("conf/conf.toml", &Cfg)
	if err != nil {
		log.Fatal("load system conf file failed:", err)
	}
}
