package mycfg

import (
	"master/utils/mycfg/cfg"
	"master/utils/mylog"
)

type CfgMgr struct {
	Data *cfg.DataConfigTable
}

var mIns *CfgMgr = nil

func Instance()*CfgMgr{
	if mIns == nil{
		mIns = &CfgMgr{}
	}
	return mIns
}

func (this *CfgMgr)InitCfg(path string)bool{
	this.Data = cfg.NewDataConfigTable()
	if err:=this.Data.Load(path);err!=nil{
		mylog.Panic(err.Error())
		return false
	}
	return true
}