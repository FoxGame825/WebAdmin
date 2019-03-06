package api

import (
		"master/define"
		"fmt"
	"time"
	"master/utils"
)


const(
	Token_Expired = time.Minute * 3
)

func CheckLoginInfo(user string,passwd string)bool{
	db:=utils.GetDbMgr().Db
	info :=&define.UserInfo{}
	if err:=db.Model(define.UserInfo{}).Where("user_name=?",user).First(info).Error;err!=nil{
		fmt.Println(err)
		return false
	}

	if info.PassWord == passwd{
		return true
	}

	return false
}

func QueryUserInfo(id int)*define.UserInfo{
	db:=utils.GetDbMgr().Db
	info :=&define.UserInfo{}
	if err:=db.Model(define.UserInfo{}).Where("id=?",id).First(info).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return info
}

func QueryAllUserInfo()[]define.UserInfo{
	infos:=make([]define.UserInfo,0)
	db:=utils.GetDbMgr().Db
	if err:=db.Model(define.UserInfo{}).Find(&infos).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return infos
}

func QueryAllPlayerInfo()[]define.PlayerInfo{
	infos:=make([]define.PlayerInfo,0)
	db:=utils.GetDbMgr().Db
	if err:=db.Table("players").Select("id,account_id, name,gold,diamond").Scan(&infos).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return infos
}

func QueryUserInfoByName(user string)*define.UserInfo{
	db:=utils.GetDbMgr().Db
	info :=&define.UserInfo{}
	if err:=db.Model(define.UserInfo{}).Where("user_name=?",user).First(info).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return info
}

func QueryUserInfoByToken(token string)*define.UserInfo{
	db:=utils.GetDbMgr().Db
	var info = &define.TokenInfo{}
	if err:=db.Model(define.TokenInfo{}).Where("token=?",token).First(info).Error;err!=nil{
		fmt.Println(err)
		return nil
	}

	return QueryUserInfo(info.UserId)
}

func MotifyUserInfo(newInfo *define.UserInfo)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(define.UserInfo{}).Updates(newInfo).Error;err!=nil{
		fmt.Println(err)
		return false
	}

	return true
}

func AddUserInfo(info *define.UserInfo)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Create(info).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func GetToken(userid int)string{
	db:=utils.GetDbMgr().Db
	var info = &define.TokenInfo{}
	if err:=db.Model(define.TokenInfo{}).Where("user_id=?",userid).First(info).Error;err!=nil{
		fmt.Println(err)
		return ""
	}
	return info.Token
}

func SetToken(userid int,token string,ip string)bool{
	old:=GetToken(userid)
	db:=utils.GetDbMgr().Db

	if !CheckTokenValid(old){
		if err:=db.Model(define.TokenInfo{}).Create(&define.TokenInfo{UserId:userid,Token:token,Ip:ip,ExpiredAt:time.Now().Add(Token_Expired)}).Error;err!=nil{
			fmt.Println(err)
			return false
		}
	}else {
		if err:=db.Model(define.TokenInfo{}).Where("user_id=?",userid).Update("token",token).Error;err!=nil{
			fmt.Println(err)
			return false
		}
	}

	return true
}

func RefreshTokenExpired(token string)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(define.TokenInfo{}).Where("token=?",token).Update("expired_at",time.Now().Add(Token_Expired)).Error;err!=nil{
		fmt.Println(err)
		return false
	}

	return true
}

func ClearTokenByUserID(userid int)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(&define.TokenInfo{}).Where("user_id=?",userid).Delete(define.TokenInfo{}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func ClearToken(token string)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(&define.TokenInfo{}).Where("token=?",token).Delete(define.TokenInfo{}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func CheckTokenValid(token string)bool{
	if len(token) <=0{
		return false
	}

	db:=utils.GetDbMgr().Db
	info:=&define.TokenInfo{}
	if err:=db.Model(&define.TokenInfo{}).Where("token=?",token).First(info).Error;err!=nil{
		fmt.Println(err)
		return false
	}

	//fmt.Println(time.Now(), info.ExpiredAt)
	return time.Now().Before(info.ExpiredAt)
}

func CheckPermission(curPermission int,checkPermission int)bool{
	return curPermission & checkPermission > 0
}

func PushLog(userid int,action string ,log string)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Create(&define.UserLogInfo{UserId:userid,Action:action,Log:log,CreatedAt:time.Now()}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func GetLog(userid int)([]define.UserLogInfo){
	infos:=make([]define.UserLogInfo,0)
	db:=utils.GetDbMgr().Db
	if err:=db.Model(define.UserLogInfo{}).Where("user_id = ?",userid).Find(&infos).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return infos
}


func CheckGoodExist(cate ,typeID int)bool{
	if cate == define.Good_Item_Type{
		if utils.GetCfgMgr().Data.ItemByID[int32(typeID)] ==nil{
			return false
		}

	}else if cate ==define.Good_Equip_Type{
		if utils.GetCfgMgr().Data.EquipByID[int32(typeID)] ==nil{
			return false
		}
	}

	return true
}

func QueryAllNoticeInfo()[]define.NoticeInfo{
	infos:=make([]define.NoticeInfo,0)
	db:=utils.GetDbMgr().Db
	if err:=db.Model(define.NoticeInfo{}).Select("*").Scan(&infos).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return infos
}


func AddNoticeInfo(title string,content string,channelId int,startTm,endTm time.Time)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Create(&define.NoticeInfo{Title:title,Content:content,ChannelId:channelId,StartTime:startTm,EndTime:endTm,CreatedAt:time.Now()}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func DelNoticeInfo(id int)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(&define.NoticeInfo{}).Where("id=?",id).Delete(&define.NoticeInfo{}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}


func AddChannelInfo(name,desc string)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Create(&define.ChannelInfo{Name:name,Desc:desc,CreatedAt:time.Now()}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}

func DelChannelInfo(id int)bool{
	db:=utils.GetDbMgr().Db
	if err:=db.Model(&define.ChannelInfo{}).Where("id=?",id).Delete(&define.ChannelInfo{}).Error;err!=nil{
		fmt.Println(err)
		return false
	}
	return true
}


func QueryAllChannelInfo()[]define.ChannelInfo{
	infos:=make([]define.ChannelInfo,0)
	db:=utils.GetDbMgr().Db
	if err:=db.Model(&define.ChannelInfo{}).Select("*").Scan(&infos).Error;err!=nil{
		fmt.Println(err)
		return nil
	}
	return infos
}



func StringToTime(tm string)time.Time{
	temp:="2006-01-02 15:04:05"
	loc,_:=time.LoadLocation("Local")
	theTm,_:=time.ParseInLocation(temp,tm,loc)
	return theTm
}
