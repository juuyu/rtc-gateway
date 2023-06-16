// Package srs
// @title
// @description
// @author njy
// @since 2023/6/1 11:17
package srs

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// TX-shanghai-20230601-001
// TX-shanghai-8e49f02caa57f13b
// TX-shanghai-1565084298
//

const (
	TENCENT = "TX"
	ALI     = "AL"
)
const (
	HANGZHOU = "hangzhou"
	SHANGHAI = "shanghai"
)

func Ping(ctx *gin.Context) {
	// 1. redis里面查询节点信息
	// 2. redis没有，数据库查
	// 3. 更新心跳
	// 4. 上下线告警，记录

	ctx.JSON(http.StatusOK, 0)
}

func FetchConf(ctx *gin.Context) {

}

func genServerCode() *string {
	return nil
}
