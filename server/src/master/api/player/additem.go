package player

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
	"master/utils/mynsq/sspb"
	"master/utils/mynsq"
	"strconv"
)

func AddItemHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	playerId,_:=strconv.Atoi(ctx.FormValue("playerId"))
	category,_:= strconv.Atoi(ctx.FormValue("category"))
	itemID,_:=strconv.Atoi(ctx.FormValue("itemId"))
	count ,_:=strconv.Atoi(ctx.FormValue("count"))

	if api.CheckTokenValid(token){
		//check data valid

		// nsq publish
		msg:=&sspb.MS2GSAddItemMsg{}
		msg.ItemCategory = int32(category)
		msg.ItemTypeID = int32(itemID)
		msg.PlayerID = int32(playerId)
		msg.Count = int32(count)
		msg.Token = token

		mynsq.Instance().Publish( uint32(sspb.WebNsqTag_AddItem),msg)

		userInfo:=api.QueryUserInfoByToken(token)
		api.PushLog(userInfo.Id,define.Action_AddItem,msg.String())

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})

	}else{
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}


}

