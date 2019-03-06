package user



import (
	"github.com/devfeel/dotweb"
				"master/api"
	"master/define"
)

//
//type userInfo struct {
//	Name string `json:"name"`
//	Avator string `json:"avator"`
//	Roles []int	`json:"roles"`
//}


func GetAllUserInfoHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")

	if api.CheckTokenValid(token){
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:api.QueryAllUserInfo()})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenInValid})
	}
}

