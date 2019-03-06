package mail

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cors"
)

func InitRoute(router dotweb.Router){
	router.POST("/mail/sendmail", SendMailHander).Use(CustomCROS())
}

func CustomCROS()dotweb.Middleware{
	option:=cors.NewConfig()
	option.SetHeader("Content-Type")
	option.SetMethod("GET,POST,OPTIONS")
	option.SetOrigin("http://localhost:8080/")
	option.Enabled()
	option.SetAllowCredentials(true)
	return cors.Middleware(option)
}
