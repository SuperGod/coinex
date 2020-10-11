package bitmex

import (
	"testing"

	. "github.com/SuperGod/trademodel"

	. "github.com/SuperGod/coinex"
	log "github.com/sirupsen/logrus"
)

func GetWSClient() (bm *BitmexWS) {
	configs, err := LoadConfigs()
	if err != nil {
		panic(err.Error())
	}
	key, secret := configs.Get("bitmextest")
	bm = NewBitmexWSTest("XBTUSD", key, secret, configs.Proxy)
	return bm
}

func TestTrade(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	tradeChan := make(chan Trade, 1024)
	clt.SetTradeChan(tradeChan)
	for trade := range tradeChan {
		log.Println(trade)
	}
}

func TestQuote(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	clt.SetSubscribe([]SubscribeInfo{SubscribeInfo{Op: BitmexWSTradeBin1m, Param: "XBTUSD"}})
	klineChan := make(chan *Candle, 1024)
	clt.SetKlineChan("1m", klineChan)
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	for kline := range klineChan {
		log.Println(kline)
	}
}

func TestWSOrder(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	subs := []SubscribeInfo{SubscribeInfo{Op: BitmexWSOrder, Param: "XBTUSD"}}
	clt.SetSubscribe(subs)
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	orderChan := make(chan []Order, 1024)
	clt.SetOrderChan(orderChan)
	for orders := range orderChan {
		log.Println("orders changed")
		for _, v := range orders {
			log.Println(v)
		}
	}
}

func TestWSPosition(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	subs := []SubscribeInfo{SubscribeInfo{Op: BitmexWSPosition, Param: "XBTUSD"}}
	clt.SetSubscribe(subs)
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	posChan := make(chan []Position, 1024)
	clt.SetPositionChan(posChan)
	for pos := range posChan {
		log.Println("position changed")
		for _, v := range pos {
			log.Println(v)
		}
	}
}

func TestWSKline(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	sub := SubscribeInfo{Op: BitmexWSTradeBin1m, Param: "XBTUSD"}
	clt.AddSubscribe(sub)
	klineChan := make(chan *Candle, 1024)
	clt.SetKlineChan("1m", klineChan)

	for candle := range klineChan {
		log.Println(candle)
	}
}

func TestWSBalance(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	clt := GetWSClient()
	subs := []SubscribeInfo{SubscribeInfo{Op: BitmexWSMargin}}
	clt.SetSubscribe(subs)
	err := clt.Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	bChan := make(chan Balance, 1024)
	clt.SetBalanceChan(bChan)
	for b := range bChan {
		log.Println("balance changed")
		log.Println(b)
	}
}
