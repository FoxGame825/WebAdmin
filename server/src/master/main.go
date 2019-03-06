package main


import (
		"github.com/devfeel/dotweb"
		"master/api"
	"master/api/user"
	"master/api/player"
	"master/api/mail"
	"master/api/notice"
	"path/filepath"
	"master/utils/mylog"
	"master/api/goods"
	"strings"
	"master/api/channel"
	"master/utils"
)


func init(){
	mylog.InitLogger(true)
}

func main() {
	mylog.Info("\n\n\n\n\n-------------startup master-----------------")

	if !utils.GetSettingMgr().InitSetting(GetPath("static/setting.json")){
		mylog.Fatal("setting load failed!!!")
	}
	mylog.Info("setting load success...")

	dbsource:=DbSourceString(utils.GetSettingMgr().MySql_Url,utils.GetSettingMgr().MySql_UserName, utils.GetSettingMgr().MySql_Password, utils.GetSettingMgr().MySql_DbName)
	if !utils.GetDbMgr().InitDB(dbsource){
		mylog.Fatal("mysql connent web db failed!!! :" + dbsource)
	}
	defer utils.GetDbMgr().DBClose()


	utils.GetCfgMgr().InitCfg(GetPath("static/DataConfig.json"))
	mylog.Info("load config success... ")


	app := dotweb.New()
	//app.SetDevelopmentMode()
	app.SetProductionMode()

	app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
		ctx.WriteString("error 404！ ", err.Error())
		ctx.End()
	})

	app.HttpServer.Router().ServerFile("/static/file/*filepath", GetPath("static/file/"))
	app.HttpServer.Renderer().SetTemplatePath(GetPath("./static/template/"))
	app.HttpServer.SetEnabledListDir(false)

	api.InitRoute(app.HttpServer.Router())
	user.InitRoute(app.HttpServer.Router())
	player.InitRoute(app.HttpServer.Router())
	mail.InitRoute(app.HttpServer.Router())
	notice.InitRoute(app.HttpServer.Router())
	goods.InitRoute(app.HttpServer.Router())
	channel.InitRoute(app.HttpServer.Router())

	err := app.StartServer(utils.GetSettingMgr().Web_Port)
	mylog.Panic(err.Error())
}


func GetPath(path string)string{
	absPath, _ := filepath.Abs(path)
	return absPath
}

//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8&parseTime=True&loc=Local"
func DbSourceString(url ,user ,passwd ,dbname string)string{
	return strings.Join([]string{user, ":", passwd, "@tcp(",url,")/", dbname, "?charset=utf8&parseTime=True&loc=Local"}, "")
}



