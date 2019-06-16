package coinex

import (
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

type DownParam interface {
	SetStart(start *int32)
	SetCount(count *int32)
	SetStartTime(startTime *time.Time)
	SetEndTime(endTime *time.Time)
}

type NewParamFunc func() DownParam

type DownFunc func(DownParam) ([]interface{}, bool, error)

type DataDownload struct {
	nDuration time.Duration
	dataCh    chan []interface{}
	paramFunc NewParamFunc
	downFunc  DownFunc
	onceCount int32
	startTime time.Time
	endTime   time.Time
	nFinish   int32 // 0 not finish, 1 finish
	nTotal    int
}

func NewDataDownload(start, end time.Time, paramFunc NewParamFunc, downFunc DownFunc, onceCount int32, nDuration time.Duration) (d *DataDownload) {
	d = new(DataDownload)
	d.paramFunc = paramFunc
	d.nDuration = nDuration
	d.dataCh = make(chan []interface{}, 1024)
	d.downFunc = downFunc
	d.onceCount = onceCount
	d.startTime = start
	d.endTime = end
	return
}

func (d *DataDownload) Start() (dataCh chan []interface{}) {
	go d.Run()
	dataCh = d.dataCh
	return
}
func (d *DataDownload) Run() {
	defer func() {
		close(d.dataCh)
		log.Debug("DataDownload finished...")
	}()

	var nStart int32
	var t1 time.Time
	var nSleep time.Duration

	for {
		params := d.paramFunc()
		params.SetStartTime(&d.startTime)
		params.SetEndTime(&d.endTime)
		params.SetCount(&d.onceCount)
		params.SetStart(&nStart)
		ret, isFinished, err := d.downFunc(params)
		if err != nil {
			return
		}
		t1 = time.Now()
		if len(ret) > 0 {
			d.dataCh <- ret
		}
		nSleep = time.Now().Sub(t1)
		if nSleep < d.nDuration {
			time.Sleep(d.nDuration - nSleep)
		}
		nStart += int32(len(ret))
		if isFinished {
			d.SetFinish(1)
			break
		}
	}
	return
}

func (d *DataDownload) SetFinish(nFinish int32) {
	atomic.StoreInt32(&d.nFinish, nFinish)
}

func (d *DataDownload) IsFinish() (bFinish bool) {
	nFinish := atomic.LoadInt32(&d.nFinish)
	bFinish = nFinish == 1
	return
}
