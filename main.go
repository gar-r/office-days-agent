package main

import (
	"fmt"
	"time"

	wifiname "git.okki.hu/garrichs/wifi-name"
	"github.com/kelseyhightower/envconfig"
)

var conf Conf

func main() {
	envconfig.MustProcess(EnvConfigPrefix, &conf)
	t := time.NewTicker(time.Duration(conf.PollIntervalSeconds) * time.Second)
	for range t.C {
		checkWifi()
	}
}

func checkWifi() {
	ssid, err := wifiname.GetSSID(conf.WifiDeviceName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ssid)
}
