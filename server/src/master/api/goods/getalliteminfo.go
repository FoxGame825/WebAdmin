package goods

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"

	"master/utils"
)



func AllItemInfoHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token := ctx.FormValue("token")
	if api.CheckTokenValid(token){
		var infos = make([]*define.GoodInfo,0)
		for _,v:= range utils.GetCfgMgr().Data.ItemByID{
			infos = append(infos,&define.GoodInfo{Id:int(v.ID),Name:v.Name,Desc:v.Desc})
		}
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed,Data:infos})
	}
	return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenInValid})
}