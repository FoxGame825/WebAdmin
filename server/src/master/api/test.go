package api

import (
	"github.com/devfeel/dotweb"
			)

func Test(ctx dotweb.Context)error{
	defer ctx.End()

	return ctx.View("login.html")

	//return ctx.WriteJson("success!!!")
}

func TestMain(ctx dotweb.Context)error{
	defer ctx.End()

	return ctx.View("main.html")
}
