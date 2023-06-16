// Package main
// @title
// @description
// @author njy
// @since 2023/5/30 11:54
package main

import (
	"log"
	"rtc-gateway/dao"
	"rtc-gateway/redis"
)

func main() {
	dao.InitMySQL()
	redis.InitRedis()

	// 启动gin
	r := StartRouter()
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("server start failed, err:", err)
	}

	// 后端路由
	//engine := StartRouter()
	//
	//// 前端路由
	//staticEngine := gin.New()
	//templateHTML, err := template.ParseFS(ui.TemplateFs, "template/**/**/*.html")
	//if err != nil {
	//	panic(err)
	//}
	//staticEngine.SetHTMLTemplate(templateHTML)
	//fads, err := fs.Sub(ui.StaticFs, "static")
	//if err != nil {
	//	panic(err)
	//}
	//staticEngine.StaticFS("/static", http.FS(fads))

	// Route
	//s.Route(c, engine)
	//s.RouteHTML(c, staticEngine)

	// 通过一个 Server 运行，判断应该走前端路由还是后端路由
	//server := &http.Server{
	//	Handler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	//		// 如果 URL 以 /static 或 /ui 开头，则走前端路由
	//		if strings.HasPrefix(request.URL.Path, "/static") ||
	//			strings.HasPrefix(request.URL.Path, "/ui") {
	//			staticEngine.ServeHTTP(writer, request)
	//			return
	//		}
	//		// 否则，走后端路由
	//		engine.ServeHTTP(writer, request)
	//	}),
	//}
	//if err = server.ListenAndServe(":8080"); err != nil {
	//	panic(err)
	//}

}
