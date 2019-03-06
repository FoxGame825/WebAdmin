package player

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
	//"master/utils/mynsq/sspb"
	//"master/utils/mynsq"
	//"strconv"
	"master/utils"
)

func AddItemHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	//playerId,_:=strconv.Atoi(ctx.FormValue("playerId"))
	//category,_:= strconv.Atoi(ctx.FormValue("category"))
	//itemID,_:=strconv.Atoi(ctx.FormValue("itemId"))
	//count ,_:=strconv.Atoi(ctx.FormValue("count"))

	if api.CheckTokenValid(token){
		//check data valid

		userInfo:=api.QueryUserInfoByToken(token)
		api.PushLog(userInfo.Id,define.Action_AddItem,ctx.Request().Form.Encode())

		utils.GetResultMgr().PushResult(token,"添加物品成功!")

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})

	}else{
		utils.GetResultMgr().PushResult(token,"添加物品失败!")
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}


}

