package bitmex

import (
	. "github.com/SuperGod/coinex"
	"github.com/spf13/viper"
)

func GetClientByViperName(cfg *viper.Viper, name string, bTest bool) (bm *Bitmex) {
	configs, err := LoadViperConfig(cfg)
	if err != nil {
		panic(err.Error())
	}
	key, secret := configs.Get(name)
	if bTest {
		bm = NewBitmexTest(key, secret)
		bm.SetDebug(true)
	} else {
		bm = NewBitmex(key, secret)
	}
	if configs.Proxy != "" {
		bm.SetProxy(configs.Proxy)
	}
	return
}

func GetClientByName(name string, bTest bool) (bm *Bitmex) {
	configs, err := LoadConfigs()
	if err != nil {
		panic(err.Error())
	}
	key, secret := configs.Get(name)
	if bTest {
		bm = NewBitmexTest(key, secret)
		bm.SetDebug(true)
	} else {
		bm = NewBitmex(key, secret)
	}
	if configs.Proxy != "" {
		bm.SetProxy(configs.Proxy)
	}
	// bm.SetDebug(true)
	return bm
}

func GetDefaultClient() *Bitmex {
	return GetClientByName("bitmex", false)
}

func GetClients(bTest bool) (clts map[string]*Bitmex) {
	configs, err := LoadConfigs()
	if err != nil {
		panic(err.Error())
	}
	clts = make(map[string]*Bitmex)
	var bm *Bitmex
	for k, v := range configs.Exchanges {
		if v.Type == "bitmex" {
			if bTest {
				bm = NewBitmexTest(v.Key, v.Secret)
				bm.SetDebug(true)
			} else {
				bm = NewBitmex(v.Key, v.Secret)
			}
			bm.SetProxy(configs.Proxy)
			clts[k] = bm
		}
	}
	return
}
