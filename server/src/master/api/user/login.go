package user

import (
	"github.com/devfeel/dotweb"
		"mytools.com/libs/uuid"
		"master/define"
	"fmt"
	"master/api"
	"master/utils/mylog"
)

type loginInfo struct {
	UserName string
	Password string
}

type responseInfo struct {
	Token string `json:"token"`
	Msg string `json:"msg"`
}

func LoginHander(ctx dotweb.Context)error{
	defer ctx.End()

	username:=ctx.PostFormValue("username")
	passwd:=ctx.PostFormValue("password")
	fmt.Println("login form data :",username,passwd)

	res :=new(responseInfo)

	if len(username) <=0 || len(passwd)<=0{
		res.Msg = "username or password format error!!"
		return ctx.WriteJson(&define.ResponseData{Code:1,Data:res})
	}

	if api.CheckLoginInfo(username,passwd){
		token:= uuid.GenUUID()
		res.Token = token
		ip:=ctx.RemoteIP()

		info:= api.QueryUserInfoByName(username)
		api.ClearTokenByUserID(info.Id)
		api.SetToken(info.Id,token,ip)

		mylog.Info(info.UserName+" login success")
		api.PushLog(info.Id,define.Action_Login,info.UserName+"登录 ip:"+ip)

		return ctx.WriteJson(&define.ResponseData{Code:0,Data:res})
	}else {
		res.Msg = "username or password error!!"
		return ctx.WriteJson(&define.ResponseData{Code:1,Data:res})
	}

}
