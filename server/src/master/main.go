package main


import (
		"github.com/devfeel/dotweb"
		"master/api"
	"master/api/user"
	"master/api/player"
	"master/api/mail"
	"master/api/notice"
	"path/filepath"
		"master/utils/mydb"
	"master/utils/mynsq"
	"master/utils/mylog"
	"master/utils/mysetting"
	"master/utils/mycfg"
	"master/api/goods"
	"strings"
	"master/api/channel"
)


func init(){
	mylog.InitLogger(true)
}

func main() {
	mylog.Info("\n\n\n\n\n-------------startup master-----------------")
	//dbMgr := utils.GetInstance()
	//if dbMgr.Init() == true {
	//	fmt.Println("mysql connent success")
	//}
	//libs.GetDBManager().Init("127.0.0.1:6379")


	if !mysetting.Instance().InitCfg(GetPath("static/setting.json")){
		mylog.Fatal("setting load failed!!!")
	}
	mylog.Info("setting load success...")

	dbsource:=DbSourceString(mysetting.Instance().MySql_Url,mysetting.Instance().MySql_UserName, mysetting.Instance().MySql_Password, mysetting.Instance().MySql_DbName)
	if !mydb.Instance().InitDB(dbsource){
		mylog.Fatal("mysql connent web db failed!!! :" + dbsource)
	}
	dbsource = DbSourceString(mysetting.Instance().GameSql_Url,mysetting.Instance().GameSql_UserName, mysetting.Instance().GameSql_Password, mysetting.Instance().GameSql_DbName)
	if !mydb.Instance().InitGameDb(dbsource){
		mylog.Fatal("mysql connent game db failed!!! :"+ dbsource)
	}
	mylog.Info("mysql connent success...")
	defer mydb.Instance().Close()

	if !mynsq.Instance().InitNsq(mysetting.Instance().Nsq_Url, mysetting.Instance().Nsq_Consumer_Topic, mysetting.Instance().Nsq_Channel){
		mylog.Fatal("nsq connent failed!!!")
	}
	mylog.Info("nsq connect success...")
	defer mynsq.Instance().Close()




	mycfg.Instance().InitCfg(GetPath("static/DataConfig.json"))
	mylog.Info("load config success... ")


	app := dotweb.New()
	//app.SetDevelopmentMode()
	app.SetProductionMode()

	app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
		ctx.WriteString("error 404！ ", err.Error())
		ctx.End()
	})

	//absPath, _ := filepath.Abs("static/file/")
	app.HttpServer.Router().ServerFile("/static/file/*filepath", GetPath("static/file/"))
	//absPath, _ = filepath.Abs("./static/template/")
	app.HttpServer.Renderer().SetTemplatePath(GetPath("./static/template/"))
	app.HttpServer.SetEnabledListDir(false)


	app.HttpServer.Router().Any("/test", api.Test)
	api.InitRoute(app.HttpServer.Router())
	user.InitRoute(app.HttpServer.Router())
	player.InitRoute(app.HttpServer.Router())
	mail.InitRoute(app.HttpServer.Router())
	notice.InitRoute(app.HttpServer.Router())
	goods.InitRoute(app.HttpServer.Router())
	channel.InitRoute(app.HttpServer.Router())


	err := app.StartServer(mysetting.Instance().Web_Port)
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



