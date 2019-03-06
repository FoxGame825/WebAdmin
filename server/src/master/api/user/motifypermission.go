package user

import (
	"github.com/devfeel/dotweb"
	"master/api"
	"master/define"
	"strconv"
	"master/utils/mylog"
	"fmt"
)

func MotifyPermissionHandler(ctx dotweb.Context)error{
	defer ctx.End()

	token :=ctx.FormValue("token")
	motifyPermission:=ctx.FormValue("motifyPermission")

	if !api.CheckTokenValid(token){
		return ctx.WriteJson(&define.ResponseData{Code:1})
	}

	userInfo:=api.QueryUserInfoByToken(token)
	if userInfo ==nil{
		return ctx.WriteJson(&define.ResponseData{Code:2})
	}
	old:= userInfo.Permission
	userInfo.Permission,_ = strconv.Atoi(motifyPermission)

	bret:=api.MotifyUserInfo(userInfo)

	if !bret{
		return ctx.WriteJson(&define.ResponseData{Code:3})
	}

	mylog.Info(userInfo.UserName + " 修改权限成功")
	lg:=fmt.Sprint("%s 修改权限成功 原权限:%s 目标权限:%s ",userInfo.UserName,old,motifyPermission)
	api.PushLog(userInfo.Id,define.Action_Motify_Permission,lg)
	return ctx.WriteJson(&define.ResponseData{Code:0})
}
