package main

import (
	"log"
	"time"

	wifiname "git.okki.hu/garrichs/wifi-name"
	"github.com/kelseyhightower/envconfig"
)

var conf Conf
var store Store

func main() {
	envconfig.MustProcess(EnvConfigPrefix, &conf)
	store = NewFileStore(conf.DataPath)
	t := time.NewTicker(time.Duration(conf.PollIntervalSeconds) * time.Second)
	for range t.C {
		if usingOfficeWifi() {
			store.Flag()
		}
	}
}

func usingOfficeWifi() bool {
	ssid, err := wifiname.GetSSID(conf.WifiDeviceName)
	if err != nil {
		log.Println(err)
		return false
	}
	return conf.WifiName == ssid
}
