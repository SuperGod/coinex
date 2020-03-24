package coinex

import (
	"strings"
	"sync/atomic"
	"time"

	"github.com/beefsack/go-rate"

	log "github.com/sirupsen/logrus"
)

type Downloader interface {
	Start() chan []interface{}
	Error() error
}

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
	lastError   error
	nMaxRetry   int
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
	d.nMaxRetry = 100
	return
}

func (d *DataDownload) Error() error {
	return d.lastError
}

func (d *DataDownload) SetMaxRetry(maxRetry int) {
	d.nMaxRetry = maxRetry
}

func (d *DataDownload) Start(err chan error) (dataCh chan []interface{}) {
	go d.Run(err)
	dataCh = d.dataCh
	return
}
func (d *DataDownload) Run(errChan chan error) {
	defer func() {
		close(d.dataCh)
		if errChan != nil {
			close(errChan)
		}
		log.Debug("DataDownload finished...")
	}()

	var nStart int32
	start := d.startTime
	var nRetry int
	for {
		d.limit.Wait()
		params := d.paramFunc()
		params.SetStartTime(&start)
		params.SetEndTime(&d.endTime)
		params.SetCount(&d.onceCount)
		params.SetStart(&nStart)
		ret, isFinished, err := d.downFunc(params)
		if err != nil {
			if nRetry < d.nMaxRetry && strings.Contains(err.Error(), "context deadline exceeded") {
				nRetry++
				log.Infof("download error:%s, retry %d time", err.Error(), nRetry)
				continue
			}
			d.lastError = err
			log.Error("download error:", params, err.Error())
			return
		}
		nRetry = 0
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
