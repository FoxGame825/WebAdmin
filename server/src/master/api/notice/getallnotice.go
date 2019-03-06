package notice

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"
)

func AllNoticeHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:= ctx.FormValue("token")
	if api.CheckTokenValid(token){
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:api.QueryAllNoticeInfo()})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}

}
