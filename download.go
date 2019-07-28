package coinex

import (
	"sync/atomic"
	"time"

	"github.com/beefsack/go-rate"

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
	nDuration   time.Duration
	binDuration time.Duration
	dataCh      chan []interface{}
	paramFunc   NewParamFunc
	downFunc    DownFunc
	onceCount   int32
	startTime   time.Time
	endTime     time.Time
	nFinish     int32 // 0 not finish, 1 finish
	nTotal      int
	limit       *rate.RateLimiter
}

func NewDataDownload(start, end time.Time, binDuration time.Duration, paramFunc NewParamFunc, downFunc DownFunc, onceCount int32, nDuration time.Duration, nLimit int) (d *DataDownload) {
	d = new(DataDownload)
	d.paramFunc = paramFunc
	d.dataCh = make(chan []interface{}, 1024)
	d.downFunc = downFunc
	d.onceCount = onceCount
	d.startTime = start
	d.endTime = end
	d.binDuration = binDuration
	d.limit = rate.New(nLimit, nDuration)
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
	start := d.startTime
	for {
		d.limit.Wait()
		params := d.paramFunc()
		params.SetStartTime(&start)
		params.SetEndTime(&d.endTime)
		params.SetCount(&d.onceCount)
		params.SetStart(&nStart)
		ret, isFinished, err := d.downFunc(params)
		if err != nil {
			return
		}
		if len(ret) > 0 {
			d.dataCh <- ret
		}

		if isFinished {
			d.SetFinish(1)
			break
		}
		if d.binDuration == 0 {
			nStart += int32(len(ret))
		} else {
			start = start.Add(d.binDuration * time.Duration(len(ret)))
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
