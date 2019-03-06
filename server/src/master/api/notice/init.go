package notice

import (
	"github.com/devfeel/dotweb"
	)

func InitRoute(router dotweb.Router){
	router.POST("/notice/sendnotice", SendNoticeHander)//.Use(CustomCROS())
	router.POST("/notice/delnotice", DelNoticeHander)//.Use(CustomCROS())
	router.GET("/notice/allnotice", AllNoticeHandler)//.Use(CustomCROS())
}
//
//func CustomCROS()dotweb.Middleware{
//	option:=cors.NewConfig()
//	option.SetHeader("Content-Type")
//	option.SetMethod("GET,POST,OPTIONS")
//	option.SetOrigin("http://localhost:8080/")
//	option.Enabled()
//	option.SetAllowCredentials(true)
//	return cors.Middleware(option)
//}
