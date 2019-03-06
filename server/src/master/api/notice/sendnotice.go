package notice

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
	"strconv"
	"strings"
		"master/utils"
)

func SendNoticeHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	title:=ctx.FormValue("noticeTitle")
	info:=ctx.FormValue("noticeContent")
	channel,_:= strconv.Atoi(ctx.FormValue("channelId"))
	startTm:=ctx.FormValue("startTime")
	endTm:=ctx.FormValue("endTime")

	if api.CheckTokenValid(token){
		startTm = strings.Replace(startTm,"T"," ",1) +":00"
		endTm = strings.Replace(endTm,"T"," ",1) +":00"

		if api.AddNoticeInfo(title,info,channel,api.StringToTime(startTm),api.StringToTime(endTm)){

			userInfo:=api.QueryUserInfoByToken(token)
			api.PushLog(userInfo.Id,define.Action_SendNotice,"title="+title +" content=" +info)

			//mynsq.Instance().PushResult(token,"添加公告成功!")
			utils.GetResultMgr().PushResult(token,"添加公告成功!")

			return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})
		}else {
			//mynsq.Instance().PushResult(token,"添加公告失败")
			utils.GetResultMgr().PushResult(token,"添加公告成功!")
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_AddNotice_Err})
		}

	}else{
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}

}


