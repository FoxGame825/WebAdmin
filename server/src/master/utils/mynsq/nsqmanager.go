package mynsq

import (
	"github.com/nsqio/go-nsq"
			"time"
	"master/utils/mylog"
	"master/utils/mynsq/sspb"
	"github.com/golang/protobuf/proto"
	"encoding/binary"
	"bytes"
		"sync"
	"fmt"
	"master/utils/mysetting"
		"strconv"
	)

type NsqMgr struct {
	Mproducer *nsq.Producer
	ProducerConfig *nsq.Config
	sync.RWMutex
}

var m_Ins *NsqMgr = nil
var m_cusumer *nsq.Consumer = nil
var mNsqQueue = make(map[string][]string,0)

func Instance()*NsqMgr{
	if m_Ins == nil{
		m_Ins = &NsqMgr{}
	}
	return m_Ins
}


type NsqData struct {
	Tag uint32
	Data []byte
}

func (this *NsqMgr)InitNsq(url ,topic,channel string)bool{
	var err error

	// consumer
	go startConsumer(url,topic,channel)
	mylog.Info("create consumer success...")

	//producer
	this.ProducerConfig = nsq.NewConfig()
	this.Mproducer,err = nsq.NewProducer(url,this.ProducerConfig)
	if err!=nil{
		panic(err)
		return false
	}
	mylog.Info("create producer success...")

	return true
}

func (this *NsqMgr)Close(){
	this.Mproducer.Stop()
	m_cusumer.Stop()
}

func (this *NsqMgr)Publish(tag uint32, pb proto.Message)bool{
	if err:=this.Mproducer.Ping();err!=nil{
		panic(err)
		return false
	}

	bt:= combineData(tag,pb)
	if err := this.Mproducer.Publish(mysetting.Instance().Nsq_Publish_Topic, bt); err != nil {
		panic("publish error: " + err.Error())
		return false
	}
	return true
}



// 消费者
func startConsumer(url ,topic,channel string) {
	var err error
	cfg:= nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second          //设置重连时间
	m_cusumer, err = nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
	// 设置消息处理函数
	m_cusumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		return decombineData(message.Body)
	}))
	// 连接到单例nsqd
	if err := m_cusumer.ConnectToNSQD(url); err != nil {
		//log.Fatal(err)
		panic(err)
	}
	//<-m_cusumer.StopChan
}


func combineData(tag uint32,pb proto.Message)[]byte{
	msg :=NsqData{}
	msg.Tag = tag

	var err error
	msg.Data,err = proto.Marshal(pb)
	if err!=nil{
		panic(err)
	}

	buf:=&bytes.Buffer{}
	err=binary.Write(buf,binary.LittleEndian,msg.Tag)
	err=binary.Write(buf,binary.LittleEndian,msg.Data)
	if err!=nil{
		panic(err)
	}
	return buf.Bytes()
}

func decombineData(data []byte)error{
	msgID:=binary.LittleEndian.Uint32(data[0:4])
	msgData:=new(NsqData)
	msgData.Tag = msgID
	msgData.Data = make([]byte,len(data)-4)
	copy(msgData.Data,data[4:])

	m_Ins.Lock()

	var err error
	var result,token string
	switch msgID {
		case uint32(sspb.WebNsqTag_AddRes):{
			pbMsg:=new(sspb.GS2MSAddPlayerMoneyRetMsg)
			err := proto.Unmarshal(msgData.Data,pbMsg)
			if err ==nil{
				if pbMsg.ErrorCode ==0{
					result = "添加资源成功"
				}else {
					result = "添加资源失败,错误码:"+strconv.FormatInt(int64(pbMsg.ErrorCode),10)
				}
				token = pbMsg.Token
			}

		}
		case uint32(sspb.WebNsqTag_AddItem):{
			pbMsg:=new(sspb.GS2MSAddItemRetMsg)
			err = proto.Unmarshal(msgData.Data,pbMsg)
			if err ==nil{

				if pbMsg.ErrorCode ==0{
					result = "添加物品成功"
				}else {
					result = "添加物品失败,错误码:"+strconv.FormatInt(int64(pbMsg.ErrorCode),10)
				}
				token = pbMsg.Token
			}
		}
		case uint32(sspb.WebNsqTag_SendMail):{
			pbMsg:=new(sspb.MS2GSSendMailRetMsg)
			err = proto.Unmarshal(msgData.Data,pbMsg)
			if err ==nil{

				if pbMsg.ErrorCode ==0{
					result = "发送邮件成功"
				}else {
					result = "发送邮件失败,错误码:"+strconv.FormatInt(int64(pbMsg.ErrorCode),10)
				}
				token = pbMsg.Token
			}
		}
		case uint32(sspb.WebNsqTag_AddNotice):{
			pbMsg:=new(sspb.VS2MSAddNoticeRetMsg)
			err = proto.Unmarshal(msgData.Data,pbMsg)
			if err ==nil{

				if pbMsg.ErrorCode ==0{
					result = "添加公告成功"
				}else {
					result = "添加公告失败,错误码:"+strconv.FormatInt(int64(pbMsg.ErrorCode),10)
				}
				token = pbMsg.Token
			}
		}
		case uint32(sspb.WebNsqTag_RemoveNotice):{
			pbMsg:=new(sspb.VS2MSRemoveNoticeRetMsg)
			err = proto.Unmarshal(msgData.Data,pbMsg)
			if err ==nil{

				if pbMsg.ErrorCode ==0{
					result = "删除公告成功"
				}else {
					result = "删除公告失败,错误码:"+strconv.FormatInt(int64(pbMsg.ErrorCode),10)
				}
				token = pbMsg.Token
			}
		}
	}

	m_Ins.PushResult(token,result)

	fmt.Println("nsq - response: ",mNsqQueue)
	m_Ins.Unlock()
	return err
}


func  (this *NsqMgr)QueryNsqResult(token string)[]string{
	this.Lock()
	val,ok:= mNsqQueue[token]
	if ok{
		delete(mNsqQueue,token)
		this.Unlock()
		return val
	}
	this.Unlock()
	return nil
}

func  (this * NsqMgr)PushResult(token,content string){
	val,ok:= mNsqQueue[token]
	if ok{
		val =append(val,content)
	}else{
		mNsqQueue[token] = make([]string,0)
		mNsqQueue[token] = append(mNsqQueue[token],content)
	}
}



