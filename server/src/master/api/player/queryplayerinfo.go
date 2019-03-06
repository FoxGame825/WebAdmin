package player

import (
	"github.com/devfeel/dotweb"
	"fmt"
	"master/define"
	)

type playerInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	LastLoginDate string `json:"last_login_date"`
	Status int `json:"status"`
	RegisteDate string `json:"registe_date"`
}

func QueryPlayerInfoHander(ctx dotweb.Context)error{
	defer ctx.End()

	token:=ctx.FormValue("token")
	tag := ctx.FormValue("searchTag")
	id:=ctx.FormValue("id")
	name:=ctx.FormValue("name")

	fmt.Println("query player info : token=",token,"tag=",tag,"id=",id,"name=",name)

	if tag== "1" { //id search

	}else { // name search

	}

	var info = new(playerInfo)
	return ctx.WriteJson(&define.ResponseData{Code:0,Data:info})

}
