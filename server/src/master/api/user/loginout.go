package user

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
	)

func LoginOutHander(ctx dotweb.Context)error{
	defer ctx.End()

	token :=ctx.FormValue("token")
	api.ClearToken(token)
	return ctx.WriteJson(&define.ResponseData{Code:0})
}
