package utils


import "sync"

type ResultMgr struct {
	ResultData map[string][]string
	sync.RWMutex
}

var mResultIns *ResultMgr = nil
func GetResultMgr()*ResultMgr{
	if mResultIns == nil{
		mResultIns = &ResultMgr{}
		mResultIns.ResultData = make(map[string][]string,0)
	}
	return mResultIns
}



func (this *ResultMgr)PushResult(token,msg string){
	this.Lock()
	val,ok:= this.ResultData[token]
	if ok{
		val =append(val,msg)
	}else{
		this.ResultData[token] = make([]string,0)
		this.ResultData[token] = append(this.ResultData[token],msg)
	}
	this.Unlock()
}

func (this *ResultMgr)PopResult(token string)[]string{
	this.Lock()
	val,ok:= this.ResultData[token]
	if ok{
		delete(this.ResultData,token)
		this.Unlock()
		return val
	}
	this.Unlock()
	return nil
}
