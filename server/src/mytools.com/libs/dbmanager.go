package libs

import (
	"bytes"
	"encoding/gob"
	"github.com/go-redis/redis"
	"log"
	"time"
	"container/list"
)

type ConPoolEvt struct {

	// 1 push 2 get
	OperCode int
	PushCon *redis.Client
	GetChan chan *redis.Client

}

type ConPool struct {
	conList *list.List
	addr string
	pwd string
	db int
	OperCh chan *ConPoolEvt
}

func (this *ConPool) Init( conSize int, addr string, pwd string, db int ) {

	this.addr = addr
	this.pwd = pwd
	this.db = db
	this.conList = list.New()

	for i := 0 ; i<conSize; i++ {
		this.conList.PushBack(redis.NewClient(&redis.Options{Addr: addr, Password:pwd, DB:db}))
	}

	this.OperCh = make(chan *ConPoolEvt)

	go this.operCon()

}

func (this *ConPool) operCon() {

	for {
		evt := <-this.OperCh
		if evt == nil {
			continue
		}
		if evt.OperCode == 1 {
			if evt.PushCon == nil {
				continue
			}
			if _, err := evt.PushCon.Ping().Result(); err != nil {
				continue
			}
			this.conList.PushBack(evt.PushCon)
		}
		if evt.OperCode == 2 {
			if evt.GetChan == nil {
				continue
			}
			if this.conList.Len() == 0 {
				evt.GetChan <- redis.NewClient(&redis.Options{Addr: this.addr, Password:this.pwd, DB:this.db})
				continue
			}
			e := this.conList.Front()
			this.conList.Remove(e)
			evt.GetChan <- e.Value.(*redis.Client)
		}
	}

}

func (this *ConPool) GetCon() *redis.Client {
	getCh := make(chan *redis.Client)
	this.OperCh <- &ConPoolEvt{2, nil, getCh}
	return <- getCh
}

func (this *ConPool) PushCon( con *redis.Client ) {
	this.OperCh <- &ConPoolEvt{1, con, nil}
}

func (this *ConPool) NewCon() {
	this.OperCh <- &ConPoolEvt{1, redis.NewClient(&redis.Options{Addr: this.addr, Password:this.pwd, DB:this.db}), nil}
}

type DBManager struct {
	conPool *ConPool
}

var dbManagerInstance *DBManager = nil

func GetDBManager() *DBManager {
	if dbManagerInstance == nil {
		dbManagerInstance = &DBManager{}
	}
	return dbManagerInstance
}

func (this *DBManager) Init(addr string) {
	this.conPool = &ConPool{}
	this.conPool.Init(50,addr,"",0)
	con := this.conPool.GetCon()
	defer this.conPool.PushCon(con)
	if _, err := con.Ping().Result(); err != nil {
		log.Println("初始化存储层失败, 连接错误 : ", err)
	}
}

func (this *DBManager) InitByDB(addr string, db int) {
	this.conPool = &ConPool{}
	this.conPool.Init(50,addr,"",db)
	con := this.conPool.GetCon()
	defer this.conPool.PushCon(con)
	if _, err := con.Ping().Result(); err != nil {
		log.Println("初始化存储层失败, 连接错误 : ", err)
	}
}

func (this *DBManager) getCon() *redis.Client {
	con := this.conPool.GetCon()

	for _, err := con.Ping().Result(); err != nil; {
		log.Println("与sql断开 重新建立连接")
		con = this.conPool.GetCon()
	}

	return con
}

func (this *DBManager) pushCon( con *redis.Client ) {
	this.conPool.PushCon(con)
}

func (this *DBManager) Exist(key string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if v, err := con.Exists(key).Result(); err != nil || v == 0 {
		return false
	}
	return true
}

func (this *DBManager) ExistByCon(key string, con *redis.Client) bool {
	if v, err := con.Exists(key).Result(); err != nil || v == 0 {
		return false
	}
	return true
}

func (this *DBManager) CheckKeyIosAdd(key string) int64 {
	con := this.getCon()
	defer this.pushCon(con)
	v, err := con.Incr(key).Result()
	if err != nil {
		return 0
	}
	return v
}

func (this *DBManager) CheckKeyIosSub(key string) int64 {
	con := this.getCon()
	defer this.pushCon(con)
	v, err := con.Decr(key).Result()
	if err != nil {
		return 0
	}
	return v
}

func (this *DBManager) GetKeyToInt(key string) int64 {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key, con) {
		return 0
	}
	v, err := con.Get(key).Int64()
	if err != nil {
		return 0
	}
	return v
}

func (this *DBManager) GetKeyToInt32(key string) int {
	return int(this.GetKeyToInt(key))
}

//返回值代表是否存在 不存在自动创建
func (this *DBManager) CheckKeyIos(key string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		con.Set(key, 1, 0)
		return false
	}
	return true
}

func (this *DBManager) DelKey(key string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return false
	}
	_, err := con.Del(key).Result()
	if err != nil {
		log.Println("删除key失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) SaveFromStruct(key string, data interface{}) bool {
	con := this.getCon()
	defer this.pushCon(con)
	var bytes bytes.Buffer
	enc := gob.NewEncoder(&bytes)
	enc.Encode(data)
	if err := con.Set(key, bytes.Bytes(), 0).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) ReadToStruct(key string, data interface{}) {
	con := this.getCon()
	defer this.pushCon(con)
	readBytes, err := con.Get(key).Bytes()
	if err != nil {
		log.Println("读取数据 : ", key, " 出错")
		return
	}
	buffer := bytes.NewBuffer(readBytes)
	dec := gob.NewDecoder(buffer)
	if dec.Decode(data) != nil {
		log.Println("解码出错")
	}
}

func (this *DBManager) SAdd(k string, v string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	_, err := con.SAdd(k, v).Result()
	if err != nil {
		log.Println("添加到集合失败 : k -> ", k, " v -> ", v)
		return false
	}
	return true
}

func (this *DBManager) SRem(k string, v string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(k,con) {
		return false
	}
	_, err := con.SRem(k, v).Result()
	if err != nil {
		log.Println("移除集合元素失败 : k -> ", k, " v -> ", v)
		return false
	}
	return true
}

func (this *DBManager) SRandMembers(key string) string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return ""
	}
	sets, err := con.SRandMember(key).Result()
	if err != nil {
		log.Println("获取集合元素失败 : ", err)
		return ""
	}
	return sets
}

func (this *DBManager) SMembers(key string) []string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return nil
	}
	sets, err := con.SMembers(key).Result()
	if err != nil {
		log.Println("获取集合元素失败 : ", err)
		return nil
	}
	return sets
}

func (this *DBManager) SIsMembers(key string, v string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return false
	}
	sets, err := con.SIsMember(key, v).Result()
	if err != nil {
		log.Println("获取集合元素失败 : ", err)
		return false
	}
	return sets
}

func (this *DBManager) SCard(key string)int64{
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return 0
	}
	sets, err := con.SCard(key).Result()
	if err != nil {
		log.Println("获取集合元素失败 : ", err)
		return 0
	}
	return sets
}

func (this *DBManager) ZAdd(key string, score float64, value string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	dbValue := redis.Z{score, value}
	_, err := con.ZAdd(key, dbValue).Result()
	if err != nil {
		log.Println("添加有序集合元素失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) ZRem(key string, score float64, value string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	dbValue := redis.Z{score, value}
	_, err := con.ZRem(key, dbValue).Result()
	if err != nil {
		log.Println("添加有序集合元素失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) ZRange(key string, min int64, max int64, rev bool) []string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return nil
	}
	var res []string = nil
	var err error = nil
	if rev {
		res, err = con.ZRevRange(key, min, max).Result()
	} else {
		res, err = con.ZRange(key, min, max).Result()
	}
	if err != nil {
		log.Println("添加有序集合元素失败 : ", err)
		return nil
	}
	return res
}

func (this *DBManager) ZCard(key string) int64 {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return 0
	}
	sets, err := con.ZCard(key).Result()
	if err != nil {
		log.Println("获取集合元素失败 : ", err)
		return 0
	}
	return sets
}

func (this *DBManager) Set(key string, data interface{}) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.Set(key, data, 0).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) SetExpiration(key string, data string, expiration time.Duration) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.Set(key, data, expiration).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) Get(key string) string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return ""
	}
	str, err := con.Get(key).Result()
	if err != nil {
		return ""
	}
	return str
}

func (this *DBManager) TTL(key string) time.Duration {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return 0
	}
	str, err := con.TTL(key).Result()
	if err != nil {
		return 0
	}
	return str
}

func (this *DBManager) LPush(key string, value string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.LPush(key, value).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

//向队列压入定时数据
func (this *DBManager) LPushExpiration(listKey string, valKey string, val string, valKeyExpiration time.Duration) bool {
	con := this.getCon()
	defer this.pushCon(con)

	if _,err := con.Set(valKey, val, valKeyExpiration).Result(); err != nil {
		log.Println("LPushExpiration 写入数据失败 : ", err)
		return false
	}
	if err := con.LPush(listKey, valKey).Err(); err != nil {
		log.Println("LPushExpiration 压入列表失败 : ", err)
		return false
	}
	return true
}

//弹出队列定时数据
func (this *DBManager) LPopExpiration(listKey string) string {
	//con := this.getCon()
	//defer this.pushCon(con)

	return this.lClearPopExpiration(listKey, true)
}

//清理队列的到期数据
func (this *DBManager) lClearPopExpiration(listKey string, isPopLastValue bool) string {
	con := this.getCon()
	defer this.pushCon(con)

	l1:
		vkey,err := con.RPop(listKey).Result()
		if err != nil || vkey == "" {
			return ""
		}
		vVal := this.Get(vkey)
		if vVal == "" {
			goto l1
		}
		if isPopLastValue {
			return vVal
		}
		con.RPush(listKey, vkey)
		return ""
}

//获取定时队列数据集
func (this *DBManager) LRangeExpiration(listKey string, count int, front bool) []string {

	con := this.getCon()
	defer this.pushCon(con)

	//清理到期数据
	this.lClearPopExpiration(listKey, false)

	var res []string
	var err error

	if front {
		 res, err = con.LRange(listKey,0,-1).Result()
	} else {
		res, err = con.LRange(listKey, -1, 0).Result()
	}

	values := make([]string,0,0)

	if err != nil {
		return values
	}

	for i,v := range res {
		if i >= count {
			continue
		}
		values = append(values, v)
	}

	return values

}

func (this *DBManager) RPush(key string, value string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.RPush(key, value).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) RPop(key string) string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return ""
	}
	res,err := con.RPop(key).Result()
	if err != nil {
		//log.Println("RPOP操作失败 : ", key, " ", err)
		return ""
	}
	return res
}

func (this *DBManager) LPop(key string) string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return ""
	}
	res,err := con.LPop(key).Result()
	if err != nil {
		//log.Println("RPOP操作失败 : ", key, " ", err)
		return ""
	}
	return res
}

func (this *DBManager) LRange(key string) []string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return nil
	}
	strArr, err := con.LRange(key, 0, 30000).Result()
	if err != nil {
		return nil
	}
	return strArr
}

func (this *DBManager) LLen( key string ) int64 {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return 0
	}
	res, err := con.LLen(key).Result()
	if err != nil {
		return 0
	}
	return res
}

func (this *DBManager) LRangeByCount(key string, maxCount int64) []string {
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return nil
	}
	strArr, err := con.LRange(key, 0, maxCount).Result()
	if err != nil {
		return nil
	}
	return strArr
}

func (this *DBManager) Keys(key string) []string {
	con := this.getCon()
	defer this.pushCon(con)
	strArr, err := con.Keys(key).Result()
	if err != nil {
		return nil
	}
	return strArr
}

func (this *DBManager) GetCon() *redis.Client {
	return this.getCon()
}

func (this *DBManager) Scan(cursor uint64, key string, count int64, client *redis.Client) ([]string,uint64) {
	con := client
	defer this.pushCon(con)
	strArr,index,err := con.Scan(cursor,key,count).Result()
	if err != nil {
		return nil,0
	}
	return strArr,index
}

func (this *DBManager) HExists(key string,field string) bool {
	con := this.getCon()
	defer this.pushCon(con)
	v,_:= con.HExists(key,field ).Result()
	return v
}

func (this *DBManager) HSet(key, field string, value interface{}) bool{
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.HSet(key,field, value).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) HGet(key, field string) string{
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return ""
	}
	str, err := con.HGet(key,field).Result()
	if err != nil {
		return ""
	}
	return str
}


func (this *DBManager) HGetAll(key string) map[string]string{
	con := this.getCon()
	defer this.pushCon(con)
	if !this.ExistByCon(key,con) {
		return map[string]string{}
	}
	str, err := con.HGetAll(key).Result()
	if err != nil {
		return map[string]string{}
	}
	return str
}

func (this *DBManager) HMSet(key string, fields map[string]interface{}) bool{
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.HMSet(key,fields).Err(); err != nil {
		log.Println("写入数据失败 : ", err)
		return false
	}
	return true
}

func (this *DBManager) HDel(key string, fields ...string) bool{
	con := this.getCon()
	defer this.pushCon(con)
	if err := con.HDel(key, fields...).Err(); err != nil {
		log.Println("删除数据失败 : ", err)
		return false
	}
	return true
}