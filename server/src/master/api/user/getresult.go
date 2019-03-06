package user

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"
		"fmt"
	"master/utils"
)

func GetResultHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	if api.CheckTokenValid(token){
		//results:=mynsq.Instance().QueryNsqResult(token)
		results:=utils.GetResultMgr().PopResult(token)

		if len(results)>0{
			fmt.Println(results)
		}
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:results})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
