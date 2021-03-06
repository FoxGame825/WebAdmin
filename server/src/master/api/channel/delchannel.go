package channel

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"
	"strconv"
		"master/utils"
)

func DelChannelInfoHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	id,_:= strconv.Atoi(ctx.FormValue("channelID"))

	if api.CheckTokenValid(token){
		api.DelChannelInfo(id)

		userInfo:=api.QueryUserInfoByToken(token)
		if !api.CheckPermission(userInfo.Permission,define.Permission_Channel_OP){
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_NO_Permission})
		}
		api.PushLog(userInfo.Id,define.Action_DelChannel,"id="+ctx.FormValue("id"))

		//mynsq.Instance().PushResult(token,"删除渠道成功!")
		utils.GetResultMgr().PushResult(token,"删除渠道成功!")

		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
