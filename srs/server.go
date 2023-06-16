// Package srs
// @title
// @description
// @author njy
// @since 2023/5/30 17:07
package srs

import (
	"fmt"
	"time"
)

func GetServer() {

	// 距离、网络、cpu、内存、空闲

	now := time.Now()
	fmt.Println(now.Unix())
}
