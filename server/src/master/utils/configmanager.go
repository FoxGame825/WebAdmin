package utils

import (
	"master/utils/cfg"
	"master/utils/mylog"
)

type CfgMgr struct {
	Data *cfg.DataConfigTable
}

var mCfgIns *CfgMgr = nil

func GetCfgMgr()*CfgMgr{
	if mCfgIns == nil{
		mCfgIns = &CfgMgr{}
	}
	return mCfgIns
}

func (this *CfgMgr)InitCfg(path string)bool{
	this.Data = cfg.NewDataConfigTable()
	if err:=this.Data.Load(path);err!=nil{
		mylog.Panic(err.Error())
		return false
	}
	return true
}