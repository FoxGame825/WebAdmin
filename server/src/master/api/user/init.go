package user


import (
	"github.com/devfeel/dotweb"
	)

func InitRoute(router dotweb.Router){
	router.POST("/user/login", LoginHander)//.Use(CustomCROS())
	router.GET("/user/logout", LoginOutHander)//.Use(CustomCROS())
	router.POST("/user/motifypermission",MotifyPermissionHandler)//.Use(CustomCROS())
	router.GET("/user/info", GetUserInfoHander)//.Use(CustomCROS())
	router.GET("/user/alluserinfo", GetAllUserInfoHander)//.Use(CustomCROS())
	router.GET("/user/heart",HeartHandler)//.Use(CustomCROS())
	router.GET("/user/getresult",GetResultHandler)//.Use(CustomCROS())
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
