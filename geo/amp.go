// Package geo
// @title
// @description
// @author njy
// @since 2023/6/2 10:02
package geo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"rtc-gateway/conf"
	"rtc-gateway/util"
)

// https://lbs.amap.com/api/webservice/guide/api/georegeo

const geoUrlFormat = "https://restapi.amap.com/v3/geocode/geo?key=%v&address=%v"

type geoResp struct {
	Status   string        `json:"status"`
	Info     string        `json:"info"`
	InfoCode string        `json:"infocode"`
	Count    string        `json:"count"`
	Geocodes []interface{} `json:"geocodes"`
}

// GetLocation 获取地址经纬度信息, 返回: lng,lat
func GetLocation(address *string) *string {
	log.Println("===开始获取地址经纬度===")
	log.Println("address:", *address)
	geoUrl := fmt.Sprintf(geoUrlFormat, conf.Cfg.Amp.GeoKey, *address)
	resp, err := http.Get(geoUrl)
	if err != nil {
		log.Fatal("获取经纬度信息失败, err:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	var geoResp geoResp
	_ = json.Unmarshal(body, &geoResp)
	if geoResp.Info != "OK" {
		log.Fatal("获取经纬度信息失败, resp:", string(body))
	}
	geocodes, _ := json.Marshal(geoResp.Geocodes[0])
	location := util.GetFieldFromJson([]string{"location"}, geocodes)
	log.Println("location:", location)
	log.Println("===获取地址经纬度成功===")
	return &location
}
