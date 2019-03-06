package user

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
)

func GetUserInfoHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	if api.CheckTokenValid(token){
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:api.QueryUserInfoByToken(token)})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenInValid})
	}
}
