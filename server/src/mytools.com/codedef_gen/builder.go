package main

import (
	"text/template"
	 "os"
	"fmt"
	"io"
	"bufio"
	"strings"
	"path/filepath"
	"log"
	"time"

)


const (
	ErrorEnumName = "CodeType"
	ErrorClassName = "CodeManager"
	ErrorPackageName = "define"
	AuthorInfo = "fox"
)


type ErrorData struct {
	ErrorCodeName string	//错误码名称
	ErrorCode string 		//错误码id
	Note string 			//注释
}

type ErrorModel struct {
	EnumName string
	ManagerName string
	PackageName string
	AuthorInfo string
	GenTime string
	Items []ErrorData
}

func main(){
	println("运行错误码生成器...")

	contents,err := loadFile(getCurPath() + "/ErrorCode.txt")

	if err !=nil {
		panic(err)
	}

	var datas []ErrorData
	datas,err =paseErrorCodeText(contents)

	if err !=nil {
		panic(err)
	}

	model := ErrorModel{}
	model.EnumName = ErrorEnumName
	model.ManagerName = ErrorClassName
	model.AuthorInfo = AuthorInfo
	model.PackageName = ErrorPackageName
	model.GenTime = time.Now().String()
	model.Items = datas

	//err = generateCSFile(model)
	//if err !=nil{
	//	panic(err)
	//}

	err = generateGOFile(model)
	if err !=nil{
		panic(err)
	}

	fmt.Println("生成成功,任意键退出...")
	fmt.Scanln()
}

//加载错误码文本
func loadFile(path string)([]string,error){
	fi,err:= os.Open(path)
	if err !=nil {
		fmt.Printf("error:%s\n",err)
		return nil,err
	}
	defer fi.Close()

	var temp []string

	br := bufio.NewReader(fi)
	for{
		a,_,c:=br.ReadLine()
		if c == io.EOF {
			break
		}
		if len(a) >0 {
			temp = append(temp,string(a))
		}
	}
	return temp,nil
}

//解析文本
func paseErrorCodeText(contents []string)([]ErrorData,error){
	println("解析错误码文本...")
	var datas []ErrorData
	for _,v := range contents{
		str := strings.Replace(v," ","",-1)
		str = strings.Replace(str,"=","|",-1)
		str = strings.Replace(str,"//","|",-1)

		sli :=strings.Split(str,"|")

		if len(sli) <3 {
			continue
		}
		//idx1 := strings.Index(str,"=")
		//idx2 := strings.Index(str,"//")

		name := sli[0]
		id := sli[1]
		content := sli[2]

		var data  = ErrorData{name,id,content}
		datas = append(datas, data)
		fmt.Println(name,id,content)
	}
	return datas , nil
}


//生成cs文件
func generateCSFile(model ErrorModel)error{
	csStr :=`
/*
	该文件由生成器自动生成,请勿手动修改!
	生成时间:{{.GenTime}}
	Author : {{.AuthorInfo}}
*/

public enum {{.EnumName}}
{
	{{range $key,$value :=.Items}}
	{{$value.ErrorCodeName}} = {{$value.ErrorCode}} , //{{$value.Note}}
	{{end}}
}


public static class {{.ManagerName}}
{
	public static string ErrorCode(int errCode)
	{
		switch(errCode)
		{
			{{range .Items}}
			case (int){{$.EnumName}}.{{.ErrorCodeName}}: return "{{.Note}}";
			{{end}}
		}
		return "没有该错误码对应信息! errCode:" + errCode;
	}
}
`


	teml,err := template.New("csTemplate").Parse(csStr)
	if err !=nil{
		panic(err)
	}


	fl, err := os.Create(getCurPath()+"/ErrorCode.cs")
	defer fl.Close()

	if err != nil{
		panic(err)
	}


	err = teml.Execute(fl,model)
	if err != nil{
		panic(err)
	}

	fl.Sync()

	return nil
}

//生成go文件
func generateGOFile(model ErrorModel)error{
	goStr:=`
/*
	该文件由生成器自动生成,请勿手动修改!
	生成时间:{{.GenTime}}
	Author : {{.AuthorInfo}}
*/
	
package {{.PackageName}}

const (
	{{range $key,$value :=.Items}}
	{{$value.ErrorCodeName}} = {{$value.ErrorCode}}  //{{$value.Note}}
	{{end}}
)

var Code_ToString = map[int]string{
	{{range $key,$value :=.Items}}
	{{$value.ErrorCodeName}} : "{{$value.Note}}", 
	{{end}}
}

`
	teml,err := template.New("goTemplate").Parse(goStr)
	if err !=nil{
		panic(err)
	}

	fl, err := os.Create(getCurPath()+"/codedef.go")
	defer fl.Close()

	if err != nil{
		panic(err)
	}


	err = teml.Execute(fl,model)
	if err != nil{
		panic(err)
	}

	fl.Sync()

	return nil
}


//获取当前执行文件路径
func getCurPath()string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}





