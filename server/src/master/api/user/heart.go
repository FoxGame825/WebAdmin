package user

import (
	"github.com/devfeel/dotweb"
	"master/define"
	"master/api"
	)

func HeartHandler(ctx dotweb.Context)error{
	defer ctx.End()
	token:= ctx.FormValue("token")

	if api.CheckTokenValid(token){
		api.RefreshTokenExpired(token)
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
