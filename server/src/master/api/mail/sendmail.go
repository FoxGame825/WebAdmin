package mail

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"strconv"
	"fmt"
	"master/api"
	//"master/utils/mynsq/sspb"
		"encoding/json"
	"master/utils/mylog"
	"master/utils"
)



func SendMailHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	title:=ctx.FormValue("mailTitle")
	mailInfo:=ctx.FormValue("mailContent")
	itemListstr:=ctx.FormValue("rewardList")
	allTag:=ctx.FormValue("toAll")
	playerID,_:= strconv.Atoi(ctx.FormValue("playerID"))
	//silver,_:= strconv.Atoi(ctx.FormValue("silver"))
	//gold,_:= strconv.Atoi(ctx.FormValue("gold"))

	fmt.Println("send mail to player : token=",token,"mailInfo=",mailInfo,"itemListstr=",itemListstr,"allTag=",allTag,"playerID=",playerID,"title=",title)


	if api.CheckTokenValid(token){
		//check data valid
		type Item struct {
			Type string
			Num string
			Cate string
		}
		
		var items []Item
		if err:=json.Unmarshal([]byte(itemListstr),&items);err!=nil{
			mylog.Warn(err.Error())
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_SendMail_Marshal_Err})
		}


		//// nsq publish
		//msg:=&sspb.MS2CSendMailMsg{}
		//msg.Topic = title
		//msg.Content = mailInfo
		//msg.Silver = int32(silver)
		//msg.Diamond = int32(gold)
		//msg.SID = 3011
		//msg.Token = token
		//if allTag == "1"{
		//	msg.IsSendAll = true
		//	msg.PlayerID = 0
		//}else{
		//	msg.PlayerID = int32(playerID)
		//	msg.IsSendAll = false
		//}
		//
		//for _,v:=range items{
		//	tp,_:= strconv.Atoi(v.Type)
		//	cate,_:= strconv.Atoi(v.Cate)
		//	num,_:= strconv.Atoi(v.Num)
		//
		//	if !api.CheckGoodExist(cate,tp){
		//		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Good_IsNot_Exist})
		//	}
		//	msg.ItemList = append(msg.ItemList,&sspb.MS2CSendMailMsg_ItemInfo{ItemTypeID:int32(tp),ItemCategory:int32(cate),Count:int32(num)})
		//}
		//mynsq.Instance().Publish( uint32(sspb.WebNsqTag_SendMail),msg)



		userInfo:=api.QueryUserInfoByToken(token)

		if !api.CheckPermission(userInfo.Permission,define.Permission_Mail_OP){
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_NO_Permission})
		}
		api.PushLog(userInfo.Id,define.Action_SendMail,ctx.Request().Form.Encode())

		utils.GetResultMgr().PushResult(token,"发送邮件成功!")

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})

	}else{
		utils.GetResultMgr().PushResult(token,"发送邮件失败!")
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}


	//var info = new(playerInfo)
	return ctx.WriteJson(&define.ResponseData{Code:0})
}

