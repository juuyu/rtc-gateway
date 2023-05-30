// Package main
// @title
// @description
// @author njy
// @since 2023/5/30 11:54
package main

import "log"

func main() {
	// 启动gin
	r := StartRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("server start failed, err:", err)
	}
}
