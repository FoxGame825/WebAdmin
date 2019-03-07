package channel

import (
		"github.com/devfeel/dotweb"
	"master/define"
	"master/api"
		"master/utils"
)

func AddChannelInfoHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	name:=ctx.FormValue("channelName")
	desc:=ctx.FormValue("channelDesc")

	if api.CheckTokenValid(token){
		api.AddChannelInfo(name,desc)

		userInfo:=api.QueryUserInfoByToken(token)
		if !api.CheckPermission(userInfo.Permission,define.Permission_Channel_OP){
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_NO_Permission})
		}
		api.PushLog(userInfo.Id,define.Action_AddChannel,"name="+name+" desc="+desc)

		//mynsq.Instance().PushResult(token,"添加渠道成功!")

		utils.GetResultMgr().PushResult(token,"添加渠道成功!")

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
