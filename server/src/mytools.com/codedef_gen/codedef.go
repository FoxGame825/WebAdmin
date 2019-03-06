
/*
	该文件由生成器自动生成,请勿手动修改!
	生成时间:2019-02-13 10:58:21.9670817 +0800 CST m=+0.090995001
	Author : fox
*/
	
package define

const (
	
	Code_Succ = 10000  //操作成功
	
	Code_LoginFail = 10001  //登录失败
	
	Code_ConMaxLimit = 10002  //并发上限
	
	Code_NoMac = 10003  //暂无机组
	
	Code_BalanceNot = 10004  //余额不足
	
	Code_ResTokenFail = 10005  //ResToken失效
	
	Code_SignFail = 10006  //Sign失效
	
	Code_ParamErr = 10007  //参数错误
	
	Code_Sending = 10008  //正在发送
	
	Code_SendFail = 10009  //发送失败
	
	Code_TargetOffline = 10010  //目标不在线
	
	Code_ExecSqlFail = 10011  //写入数据库失败
	
	Code_UserAccoutIsNil = 20000  //用户账号为空
	
	Code_UserAccoutIsExist = 20001  //用户账户已经存
	
	Code_PassWdLen = 20002  //用户密码长度不符合标准
	
	Code_RepassWdIsDiffPassWd = 20003  //两次密码不相同
	
	Code_QQIsNil = 20004  //QQ账号为空
	
	Code_EmailIsNil = 20005  //Email账号为空
	
	Code_UserNameIsNil = 20006  //用户名为空
	
	Code_AlipayIsNil = 20007  //支付宝账号为空
	
	Code_SafeCodeIsNil = 20008  //安全码为空
	
	Code_UserAccoutIsNotExist = 20009  //用户账号不存在
	
	Code_UserInfoIsErr = 20010  //用户数据出错
	
	Code_UserOldPwdIsErr = 20011  //旧密码不正确
	
	Code_PassWdIsErr = 20012  //密码不正确
	
	Code_UserSignNumIsMax = 20013  //用户签名数已经达到最大了
	
	CardCore_ParamFail = 30001  //参数错误
	
	CardCore_Expire = 30002  //手机已经过期
	
	CardCore_ChannelNonExistent = 30003  //专属通道不存在
	
	CardCore_NodeKeyFail = 30004  //帐号节点未注册
	
	Code_ProjectNameIsNil = 40000  //产品名称为空
	
	Code_ProjectCardAccPriceParseErr = 40002  //卡商价格解析错误
	
	Code_ProjectDevlopAccPriceParseErr = 40003  //开发者分成价格解析错误
	
	Code_ProjectPlatformPriceParseErr = 40004  //平台价格解析错误
	
	Code_ProjectRemainParseErr = 40005  //库存解析错误
	
	Code_ProjectNotExist = 40006  //产品不存在
	
	Code_BlcokUserExist = 40007  //封号id已存在
	
	Code_ProjectPageArgsError = 40008  //获取项目参数解析错误
	
	Code_Score_Info_Accout_NotExist = 50000  //积分用户不存在
	
	Code_Score_Info_CardAcc_NotExist = 50001  //积分卡商用户不存在
	
	Code_Score_Info_DevlopAcc_NotExist = 50002  //积分开发者不存在
	
)

var Code_ToString = map[int]string{
	
	Code_Succ : "操作成功", 
	
	Code_LoginFail : "登录失败", 
	
	Code_ConMaxLimit : "并发上限", 
	
	Code_NoMac : "暂无机组", 
	
	Code_BalanceNot : "余额不足", 
	
	Code_ResTokenFail : "ResToken失效", 
	
	Code_SignFail : "Sign失效", 
	
	Code_ParamErr : "参数错误", 
	
	Code_Sending : "正在发送", 
	
	Code_SendFail : "发送失败", 
	
	Code_TargetOffline : "目标不在线", 
	
	Code_ExecSqlFail : "写入数据库失败", 
	
	Code_UserAccoutIsNil : "用户账号为空", 
	
	Code_UserAccoutIsExist : "用户账户已经存", 
	
	Code_PassWdLen : "用户密码长度不符合标准", 
	
	Code_RepassWdIsDiffPassWd : "两次密码不相同", 
	
	Code_QQIsNil : "QQ账号为空", 
	
	Code_EmailIsNil : "Email账号为空", 
	
	Code_UserNameIsNil : "用户名为空", 
	
	Code_AlipayIsNil : "支付宝账号为空", 
	
	Code_SafeCodeIsNil : "安全码为空", 
	
	Code_UserAccoutIsNotExist : "用户账号不存在", 
	
	Code_UserInfoIsErr : "用户数据出错", 
	
	Code_UserOldPwdIsErr : "旧密码不正确", 
	
	Code_PassWdIsErr : "密码不正确", 
	
	Code_UserSignNumIsMax : "用户签名数已经达到最大了", 
	
	CardCore_ParamFail : "参数错误", 
	
	CardCore_Expire : "手机已经过期", 
	
	CardCore_ChannelNonExistent : "专属通道不存在", 
	
	CardCore_NodeKeyFail : "帐号节点未注册", 
	
	Code_ProjectNameIsNil : "产品名称为空", 
	
	Code_ProjectCardAccPriceParseErr : "卡商价格解析错误", 
	
	Code_ProjectDevlopAccPriceParseErr : "开发者分成价格解析错误", 
	
	Code_ProjectPlatformPriceParseErr : "平台价格解析错误", 
	
	Code_ProjectRemainParseErr : "库存解析错误", 
	
	Code_ProjectNotExist : "产品不存在", 
	
	Code_BlcokUserExist : "封号id已存在", 
	
	Code_ProjectPageArgsError : "获取项目参数解析错误", 
	
	Code_Score_Info_Accout_NotExist : "积分用户不存在", 
	
	Code_Score_Info_CardAcc_NotExist : "积分卡商用户不存在", 
	
	Code_Score_Info_DevlopAcc_NotExist : "积分开发者不存在", 
	
}

