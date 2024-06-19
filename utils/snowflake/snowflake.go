package snowflake

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	sequenceBits  = int64(12)                               //序列id位数
	maxSequenceID = int64(-1) ^ (int64(-1) << sequenceBits) //最大序列id
	timeLeft      = uint8(22)                               //时间id向左移位的量
	machineLeft   = uint8(17)                               //机器id向左移位的量
	serviceLeft   = uint8(12)                               //服务id向左移位的量
	twepoch       = int64(1667972427000)                    //初始毫秒,时间是: Wed Nov  9 13:40:27 CST 2022
)

type Worker struct {
	sync.Mutex
	lastStamp  int64
	machineID  int64 //机器id,0~31
	serviceID  int64 //服务id,0~31
	sequenceID int64
}

var count int32

func NewWorker(machineID, serviceID int64) *Worker {
	return &Worker{
		lastStamp:  0,
		machineID:  machineID,
		serviceID:  serviceID,
		sequenceID: 0,
	}
}

func (w *Worker) getID() int64 {
	w.Lock()
	defer w.Unlock()
	mill := time.Now().UnixMilli()
	if mill == w.lastStamp {
		w.sequenceID = (w.sequenceID + 1) & maxSequenceID
		if w.sequenceID == 0 {
			for mill > w.lastStamp {
				mill = time.Now().UnixMilli()
			}
		}
	} else {
		w.sequenceID = 0
	}
	w.lastStamp = mill
	id := (w.lastStamp-twepoch)<<timeLeft | w.machineID<<machineLeft | w.serviceID<<serviceLeft | w.sequenceID
	return id
}

func GetResp() string {

	count1 := atomic.AddInt32(&count, 1)

	var work = NewWorker(963, 963)
	prefix := work.getID()
	if count1 > 9998 {
		atomic.StoreInt32(&count, 1)
	}
	orderNumber := fmt.Sprintf("%d%04d", prefix, count)

	return orderNumber
}
