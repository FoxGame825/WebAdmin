package channel



import (
	"github.com/devfeel/dotweb"
	)

func InitRoute(router dotweb.Router){
	router.GET("/channel/allchannel", AllChannelInfoHandler)//.Use(CustomCROS())
	router.POST("/channel/addchannel", AddChannelInfoHandler)//.Use(CustomCROS())
	router.POST("/channel/delchannel", DelChannelInfoHandler)//.Use(CustomCROS())
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
