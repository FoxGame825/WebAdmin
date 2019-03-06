package player

import (
	"github.com/devfeel/dotweb"
	"master/api"
	//"strconv"
	"master/define"
	//"master/utils/mynsq/sspb"
	//	"master/utils/mynsq"
	"master/utils"
)

func AddResHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	//playerId,_:= strconv.Atoi(ctx.FormValue("playerID"))
	//currency,_:= strconv.Atoi(ctx.FormValue("currency"))
	//count,_ := strconv.Atoi(ctx.FormValue("count"))

	if api.CheckTokenValid(token){
		//check data valid

		//// nsq publish
		//msg:=&sspb.MS2GSAddPlayerMoneyMsg{}
		//msg.Currency = int32(currency)
		//msg.Value = int32(count)
		//msg.PlayerID = int32(playerId)
		//msg.Token = token
		//
		//mynsq.Instance().Publish( uint32(sspb.WebNsqTag_AddRes),msg)

		userInfo:=api.QueryUserInfoByToken(token)
		api.PushLog(userInfo.Id,define.Action_AddRes,ctx.Request().Form.Encode())

		utils.GetResultMgr().PushResult(token,"添加资源成功!")

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})

	}else{
		utils.GetResultMgr().PushResult(token,"添加资源失败!")
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}

}
