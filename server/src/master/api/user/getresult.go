package user

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"
	"master/utils/mynsq"
	"fmt"
)

func GetResultHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	if api.CheckTokenValid(token){
		results:=mynsq.Instance().QueryNsqResult(token)
		if len(results)>0{
			fmt.Println(results)
		}
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:results})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
