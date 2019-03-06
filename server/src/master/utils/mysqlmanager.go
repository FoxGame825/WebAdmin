package utils

import "database/sql"
import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"io/ioutil"
	"path/filepath"
	"master/define"
	"time"
	)

type MySqlManager struct {
	Db *sql.DB
}

var mysqlInstance *MySqlManager = nil
var autoCreateTable bool = false

func GetInstance() *MySqlManager{
	if mysqlInstance == nil{
		mysqlInstance = &MySqlManager{}
	}
	return mysqlInstance
}

func (this *MySqlManager)Init()bool{
	db,err:= sql.Open("mysql","root:@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		fmt.Println(err)
		return false
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err = db.Ping(); err !=nil{
		log.Println(err)
		return false
	}

	this.Db = db

	createTable(db)

	return true
}

func (this *MySqlManager)Close(){
	this.Db.Close()
}

func (this *MySqlManager)Ping()bool{
	return this.Db.Ping() !=nil
}

//创建表
func createTable(db *sql.DB){
	absPath, _ := filepath.Abs("./src/master/sql/create_admin_userinfo_table.sql")
	sqlBytes, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Println(err)
		return
	}
	sqlTable1 := string(sqlBytes)

	absPath, _ = filepath.Abs("./src/master/sql/create_admin_userlog_table.sql")
	sqlBytes,err =ioutil.ReadFile(absPath)
	if err != nil {
		log.Println(err)
		return
	}
	sqlTable2 := string(sqlBytes)

	absPath, _ = filepath.Abs("./src/master/sql/create_admin_usertoken_table.sql")
	sqlBytes,err =ioutil.ReadFile(absPath)
	if err != nil {
		log.Println(err)
		return
	}
	sqlTable3 := string(sqlBytes)

	_,err =db.Exec(sqlTable1)
	if err!=nil{
		log.Println(err)
		return
	}

	_,err =db.Exec(sqlTable2)
	if err!=nil{
		log.Println(err)
		return
	}

	_,err =db.Exec(sqlTable3)
	if err!=nil{
		log.Println(err)
		return
	}

	fmt.Println("createTable ok!")
}


//添加后台用户
func (this *MySqlManager)AddUserInfo(info *define.UserInfo) bool{
	if this.Ping(){
		return false
	}

	timeStr:=time.Now().Format("2006-01-02 15:04:05")

	rs,err:=this.Db.Exec("INSERT INTO admin_userinfo(username,password,permission,created) VALUES(?,?,?,?)",info.UserName,info.PassWord,info.Permission,timeStr)
	if err!=nil{
		log.Println(err)
		return false
	}

	id,err :=rs.LastInsertId()
	if err!=nil{
		log.Println(err)
		return false
	}
	fmt.Println("add user id:",id)

	return true
}


//删除后台用户
func (this *MySqlManager)DelUserInfo(id int)bool{
	if this.Ping(){
		return false
	}

	rs,err:=this.Db.Exec("DELETE FROM admin_userinfo WHERE id=?",id)
	if err!=nil{
		log.Println(err)
		return false
	}

	_,err =rs.RowsAffected()
	if err !=nil{
		log.Println(err)
		return false
	}

	return true
}

//修改后台用户
func (this *MySqlManager)UpdateUserInfo(info *define.UserInfo)bool{
	if this.Ping(){
		return false
	}

	stmt,err:=this.Db.Prepare("UPDATE admin_userinfo SET password=?,permission=?,avator=? WHERE id=?")
	defer stmt.Close()
	if err!=nil{
		log.Println(err)
		return false
	}

	rs,err:=stmt.Exec(info.PassWord,info.Permission,info.Avator)
	if err!=nil{
		log.Println(err)
		return false
	}

	_,err =rs.RowsAffected()
	if err!=nil{
		log.Println(err)
		return false
	}

	return true
}


//查询后台用户信息
func (this *MySqlManager)QueryUserInfo(id int)*define.UserInfo{
	if this.Ping(){
		return nil
	}

	var info = new(define.UserInfo)
	err:= this.Db.QueryRow("SELECT id,username,password,permission,created from admin_userinfo WHERE id=?",id).Scan(
		&info.Id,&info.UserName,&info.PassWord,&info.Permission,&info.CreatedAt)
	if err!=nil{
		log.Println(err)
		return nil
	}

	return info
}
func (this *MySqlManager)QueryUserInfoByName(user string)*define.UserInfo{
	if this.Ping(){
		return nil
	}

	var info = new(define.UserInfo)
	err:= this.Db.QueryRow("SELECT id,username,password,permission,created from admin_userinfo WHERE username=?",user).Scan(
		&info.Id,&info.UserName,&info.PassWord,&info.Permission,&info.CreatedAt)
	if err!=nil{
		log.Println(err)
		return nil
	}

	return info
}

func (this *MySqlManager)QueryUserTokenInfoByToken(token string)*define.TokenInfo{
	if this.Ping(){
		return nil
	}

	var info = new(define.TokenInfo)
	err:= this.Db.QueryRow("SELECT userid,token,ip,created from admin_usertoken WHERE token=?",token).Scan(
		&info.UserId,&info.Token,&info.Ip,&info.ExpiredAt)
	if err!=nil{
		log.Println(err)
		return nil
	}

	return info
}

func (this *MySqlManager)QueryAllUserInfo()[]*define.UserInfo{
	if this.Ping(){
		return nil
	}

	rows,err:=this.Db.Query("SELECT id,username,password,permission,created FROM admin_userinfo ")
	defer rows.Close()

	if err!=nil{
		log.Println(err)
		return nil
	}

	infos:=make([]*define.UserInfo,0)
	for rows.Next(){
		var info define.UserInfo
		rows.Scan(&info.Id,&info.UserName,&info.PassWord,&info.Permission,&info.CreatedAt)
		infos = append(infos,&info)
	}

	if err = rows.Err();err!=nil{
		log.Println(err)
		return nil
	}

	return infos
}



//记录用户后台操作
func (this *MySqlManager)PushUserLog(userid int,action int,logStr string){
	if this.Ping(){
		return
	}

	timeStr:=time.Now().Format("2006-01-02 15:04:05")

	rs,err:=this.Db.Exec("INSERT INTO admin_userlog(userid,action,log,created) VALUES(?,?,?,?)",userid,action,logStr,timeStr)
	if err!=nil{
		log.Println(err)
		return
	}

	id,err :=rs.LastInsertId()
	if err!=nil{
		log.Println(err)
		return
	}
	fmt.Println("push log id:",id," action:",action,"log:",logStr)
}

//获取用户后台操作记录
func (this *MySqlManager)QueryUserLog(userid int)([]*define.UserLogInfo){
	if this.Ping(){
		return nil
	}

	rows,err:=this.Db.Query("SELECT id,userid,action,log,created FROM admin_userlog WHERE userid=?",userid)
	defer rows.Close()

	if err!=nil{
		log.Println(err)
		return nil
	}

	infos:=make([]*define.UserLogInfo,0)
	for rows.Next(){
		var info define.UserLogInfo
		rows.Scan(&info.Id,&info.UserId,&info.Action,&info.Log,&info.CreatedAt)
		infos = append(infos,&info)
	}

	if err = rows.Err();err!=nil{
		log.Println(err)
		return nil
	}

	return infos
}

//添加token
func (this *MySqlManager)SetToken(userid int,token string)bool{
	if this.Ping(){
		return false
	}

	timeStr:=time.Now().Format("2006-01-02 15:04:05")
	rows,err:=this.Db.Query("INSERT INTO admin_usertoken(userid,token,created) VALUES(?,?,?) ON DUPLICATE KEY UPDATE token=VALUES (?)",userid,token,timeStr,token)
	defer rows.Close()

	if err!=nil{
		log.Println(err)
		return false
	}


	return true
}

//获取token
func (this *MySqlManager)QueryToken(userid int)string{
	if this.Ping(){
		return ""
	}

	var token string
	err:=this.Db.QueryRow("SELECT token FROM admin_userlog WHERE userid=?",userid).Scan(&token)
	if err!=nil{
		log.Println(nil)
		return ""
	}

	return token
}

//删除token
func (this *MySqlManager)ClearToken(userid int)bool{
	if this.Ping(){
		return false
	}

	rs,err:=this.Db.Exec("DELETE FROM admin_userlog WHERE userid=?",userid)
	if err!=nil{
		log.Println(err)
		return false
	}

	_,err =rs.RowsAffected()
	if err !=nil{
		log.Println(err)
		return false
	}


	return true
}
