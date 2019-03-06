package player

import (
	"github.com/devfeel/dotweb"
		)

func InitRoute(router dotweb.Router){
	router.GET("/player/allplayerinfo", AllPlayerInfoHandler)//.Use(CustomCROS())
	router.POST("/player/additem", AddItemHander)//.Use(CustomCROS())
	router.POST("/player/addres", AddResHandler)//.Use(CustomCROS())
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
