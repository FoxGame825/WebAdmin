package mylog

import (
	"log"
	"time"
	"os"
)


var writeLogger *log.Logger = nil

func InitLogger(tofile bool){
	if tofile{
		file:="./"+time.Now().Format("2006-01-02") +"_log.txt"
		logfl,err:=os.OpenFile(file,os.O_CREATE | os.O_APPEND|os.O_RDWR,0766)
		if err!=nil{
			panic(err)
		}
		writeLogger = log.New(logfl,"",log.LstdFlags)
	}
}


func Info(content string) {
	str:="[Info] "+content
	log.Println(str)
	if writeLogger !=nil{
		writeLogger.Println(str)
	}
}

func Warn(content string) {
	str:="[Warn] "+content
	log.Println(str)
	if writeLogger !=nil{
		writeLogger.Println(str)
	}
}

func Fatal(content string) {
	str:="[Fatal] "+content
	log.Println(str)
	if writeLogger !=nil{
		writeLogger.Fatalln(str)
	}
}

func Panic(content string) {
	str:="[Panic] "+content
	log.Println(str)
	if writeLogger !=nil{
		writeLogger.Panicln(str)
	}
}


