package utils

import (
	"github.com/jinzhu/gorm"
		"master/define"
	"master/utils/mylog"
 _ "github.com/go-sql-driver/mysql"
)

type DbMgr struct {
	Db *gorm.DB
	//GameDb *gorm.DB
}

var mDbIns *DbMgr = nil

func GetDbMgr()*DbMgr{
	if mDbIns ==nil{
		mDbIns = &DbMgr{}
	}
	return mDbIns
}

func (this *DbMgr)InitDB(source string)bool{
	var err error
	this.Db, err = gorm.Open("mysql",  source)
	if err != nil {
		panic(err)
		return false
	}

	if !this.Db.HasTable(&define.UserInfo{}){
		if err:=this.Db.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&define.UserInfo{}).Error;err!=nil{
			panic(err)
			return false
		}
		mylog.Info("create table :"+ define.UserInfo{}.TableName())
	}
	if !this.Db.HasTable(&define.UserLogInfo{}){
		if err:=this.Db.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&define.UserLogInfo{}).Error;err!=nil{
			panic(err)
			return false
		}
		mylog.Info("create table :"+ define.UserLogInfo{}.TableName())
	}
	if !this.Db.HasTable(&define.TokenInfo{}){
		if err:= this.Db.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&define.TokenInfo{}).Error;err!=nil{
			panic(err)
			return false
		}
		mylog.Info("create table :"+ define.TokenInfo{}.TableName())
	}
	if !this.Db.HasTable(&define.NoticeInfo{}){
		if err:= this.Db.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&define.NoticeInfo{}).Error;err!=nil{
			panic(err)
			return false
		}
		mylog.Info("create table :"+ define.NoticeInfo{}.TableName())
	}
	if !this.Db.HasTable(&define.ChannelInfo{}){
		if err:= this.Db.Set("gorm:table_options","ENGINE=InnoDB").CreateTable(&define.ChannelInfo{}).Error;err!=nil{
			panic(err)
			return false
		}
		mylog.Info("create table :"+ define.ChannelInfo{}.TableName())
	}

	this.Db.DB().SetMaxIdleConns(10)
	this.Db.DB().SetMaxOpenConns(100)

	return true
}
//
//func (this *DbMgr)InitGameDb(source string)bool{
//	var err error
//	this.GameDb, err = gorm.Open("mysql", source)
//	if err != nil {
//		panic(err)
//		return false
//	}
//	this.GameDb.DB().SetMaxIdleConns(10)
//	this.GameDb.DB().SetMaxOpenConns(100)
//	return true
//}


func (this *DbMgr)DBClose(){
	this.Db.Close()
	//this.GameDb.Close()
}

