package libs

import (
	"fmt"
	"github.com/devfeel/dotweb"
)

func InitWebApi(port int, initRouteFun func(server *dotweb.HttpServer)) {
	app := dotweb.New()
	//app.SetDevelopmentMode()
	app.SetProductionMode()
	//设置自定义异常处理接口
	app.SetExceptionHandle(func(ctx dotweb.Context, err error) {
		ctx.WriteString("error 404！ ", err.Error())
		ctx.End()
	})
	initRouteFun(app.HttpServer)
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}
