package define

import (
	"time"
	)

const(
	Permission_AddItem_OP = 1<<0	//给玩家添加物品,装备,银两,元宝权限
	Permission_Mail_OP = 1<<1		//给玩家发送邮件权限
	Permission_Notice_OP = 1<<2		//添加删除公告权限
	Permission_Motify_Permission_OP = 1<<3	//修改其他用户权限
	Permission_Channel_OP = 1<<4 //渠道操作权限
)
func PermissionString(per int)string{
	switch per {
	case Permission_AddItem_OP:return "给玩家添加物品,装备,银两,元宝权限"
	case Permission_Mail_OP:return "给玩家发送邮件权限"
	case Permission_Notice_OP:return "添加删除公告权限"
	case Permission_Motify_Permission_OP:return "修改其他用户权限"
	case Permission_Channel_OP:return "渠道操作权限"
	default:
		return ""
	}
}

const(
	Action_Motify_Permission = "修改权限"	//修改权限操作
	Action_AddItem = "添加道具"	//添加物品操作
	Action_AddRes = "添加资源" //添加资源操作
	Action_SendMail = "发送邮件" //发送邮件操作
	Action_SendNotice = "添加公告" //发送公告操作
	Action_DeleteNotice = "删除公告" //删除公告操作
	Action_Login = "登录"	//登录操作
	Action_AddChannel = "添加渠道"
	Action_DelChannel = "删除渠道"
)
func ActionString(action string)string{
	switch action {
	case Action_Motify_Permission:return "修改权限"
	case Action_AddItem:return "添加道具"
	case Action_AddRes:return "添加资源"
	case Action_SendMail:return "发送邮件"
	case Action_SendNotice:return "添加公告"
	case Action_DeleteNotice:return "删除公告"
	case Action_Login:return "登录"
	case Action_AddChannel:return "添加渠道"
	case Action_DelChannel:return "删除渠道"
	default:
		return ""
	}
}

const(
	Code_Successed = 0
	Code_TokenExpired = 1000	//token过期
	Code_TokenInValid = 1001	//token无效
	Code_Protobuf_Marshal_Err = 2000 //protobuf 序列化错误
	Code_SendMail_Marshal_Err = 2001 //发送邮件参数序列化错误
	Code_Good_IsNot_Exist = 2003 //物品不存在
	Code_AddNotice_Err = 2004	//添加公告失败
	Code_NO_Permission = 2005	//权限不足
	Code_User_Not_Exist = 2006 //用户不存在
)
func CodeString(code int)string{
	switch code {
	case Code_Successed:return "成功"
	case Code_TokenExpired:return "token过期"
	case Code_TokenInValid:return "token无效"
	case Code_Protobuf_Marshal_Err:return "protobuf 序列化错误"
	case Code_SendMail_Marshal_Err:return "发送邮件参数序列化错误"
	case Code_Good_IsNot_Exist:return "物品不存在"
	case Code_AddNotice_Err:return "添加公告失败"
	case Code_NO_Permission:return "权限不足"
	case Code_User_Not_Exist:return "用户不存在"
	default:
		return ""
	}
}

const(
	Good_Item_Type = 1
	Good_Equip_Type = 2
)

type UserInfo struct {
	Id int	`json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserName string	`json:"username" form:"username" gorm:"type:varchar(30);not null;unique"`
	PassWord string	`json:"password" form:"password" gorm:"type:varchar(40)"`
	Permission int `json:"permission" form:"permission"`
	Avator string `json:"avator" form:"avator"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
func (UserInfo)TableName()string{return "admin_userinfo"}

type UserLogInfo struct {
	Id int `json:"id" form:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId int `json:"user_id" form:"userid" gorm:"not null"`
	Action string `json:"action" form:"action" gorm:"not null"`
	Log string `json:"log" form:"log" gorm:"type:varchar(2048)"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
func (UserLogInfo)TableName()string{return "admin_userloginfo"}


type TokenInfo struct {
	UserId int `json:"user_id" form:"userid" gorm:"primary_key"`
	Token string `json:"token" form:"token" gorm:"not null"`
	Ip string `json:"ip" form:"ip"`
	ExpiredAt time.Time `json:"expired_at" form:"expired_at"`
}
func (TokenInfo)TableName()string{return "admin_tokeninfo"}

type NoticeInfo struct {
	Id int `json:"id" gorm:"primary_key`
	ChannelId int `json:"channel_id"`
	Title string `json:"title" gorm:not null`
	Content string `json:"content" `
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`	
}
func (NoticeInfo)TableName()string{return "admin_noticeinfo"}

type ChannelInfo struct {
	Id int `json:"id" gorm:"primary_key`
	Name string `json:"name"`
	Desc string `json:"desc"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}
func (ChannelInfo)TableName()string{return "admin_channelinfo"}



type PlayerInfo struct {
	Id int `json:"id"`
	AccountId int `json:"account_id"`
	Name string `json:"name"`
	Gold int64 `json:"gold"`
	Diamond int64 `json:"diamond"`
	LockStatus int `json:"lock_status"`
}


type GoodInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}


type ResponseData struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}


func CodeToString(code int)string{
	return ""
}


