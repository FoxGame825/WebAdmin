package api

import (
	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/cors"
	)


func InitRoute(router dotweb.Router){
	router.GET("/login", LoginViewHander).Use(CustomCROS())
	router.GET("/main", MainViewHander).Use(CustomCROS())
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


func LoginViewHander(ctx dotweb.Context)error{
	defer ctx.End()
	return ctx.View("bootstrap_login.html")
}


func MainViewHander(ctx dotweb.Context)error{
	defer ctx.End()
	return ctx.View("bootstrap_main.html")
}


