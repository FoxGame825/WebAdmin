package uid

import (
	"sync"
	"time"
)

const (
	startTime   = 1514736000 // 2018/01/01 00:00:00
	maxIndex    = 1<<20 - 1
	maxServerID = 1<<14 - 1
)

var (
	lastTime  = time.Now().Unix() - startTime
	lastIndex int64
	nodeID    int64
	lock      sync.Mutex
)

// 每个服务器的serverId不一样
func Init(serverId int) {
	nodeID = int64(serverId)
}

// 生成唯一id
func Gen() int64 {
	if nodeID == 0 {
		panic("请先调用 Init(nodeId int)   nodeId")
	}
	lock.Lock()
	defer lock.Unlock()
	for {
		t := time.Now().Unix() - startTime
		if t == lastTime {
			if lastIndex >= maxIndex {
				continue
			}

			lastIndex++
		} else {
			lastIndex = 0
			lastTime = t
		}
		id := (lastTime<<34 | (nodeID << 20) | lastIndex) & (1<<63 - 1)
		return id
	}
}
