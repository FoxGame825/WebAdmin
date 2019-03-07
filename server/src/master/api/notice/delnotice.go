package notice

import (
	"github.com/devfeel/dotweb"
		"master/define"
	"master/api"
		"strconv"
		"master/utils"
)

func DelNoticeHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	id,_:= strconv.Atoi(ctx.FormValue("noticeID"))

	if api.CheckTokenValid(token){
		api.DelNoticeInfo(id)

		userInfo:=api.QueryUserInfoByToken(token)

		if !api.CheckPermission(userInfo.Permission,define.Permission_Notice_OP){
			return ctx.WriteJson(&define.ResponseData{Code:define.Code_NO_Permission})
		}

		api.PushLog(userInfo.Id,define.Action_DeleteNotice,"noticeID="+ctx.FormValue("noticeID"))
		//mynsq.Instance().PushResult(token,"删除公告成功!")
		utils.GetResultMgr().PushResult(token,"删除渠道成功!")
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_Successed})
	}else {
		return ctx.WriteJson(&define.ResponseData{Code:define.Code_TokenExpired})
	}
}
