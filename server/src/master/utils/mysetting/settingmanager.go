package mysetting

import (
	"os"
	"encoding/json"
	"fmt"
)

type SettingMgr struct {
	MySql_Url string `json:"mysql_url"`
	MySql_UserName string `json:"mysql_username"`
	MySql_Password string `json:"mysql_password"`
	MySql_DbName string `json:"mysql_dbname"`
	
	GameSql_Url string `json:"game_sql_url"` 
	GameSql_UserName string `json:"game_sql_user_name"`
	GameSql_Password string `json:"game_sql_password"`
	GameSql_DbName string `json:"game_sql_db_name"`

	Nsq_Url string	`json:"nsq_url"`
	Nsq_Consumer_Topic string `json:"nsq_consumer_topic"`
	Nsq_Publish_Topic string `json:"nsq_publish_topic"`
	Nsq_Channel string `json:"nsq_channel"`

	Web_Port int `json:"web_port"`
}

var mIns *SettingMgr =nil

func Instance()*SettingMgr{
	if mIns == nil{
		mIns = &SettingMgr{}
	}
	return mIns
}

func (this *SettingMgr)InitCfg(path string)bool{
	fl,_:=os.Open(path)
	defer fl.Close()

	decoder:=json.NewDecoder(fl)
	err:=decoder.Decode(this)
	if err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}
