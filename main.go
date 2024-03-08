package main

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	wifiname "github.com/yelinaung/wifi-name"
)

func main() {
	var conf Conf
	envconfig.MustProcess(EnvConfigPrefix, &conf)
	t := time.NewTicker(time.Duration(conf.PollIntervalSeconds) * time.Second)
	for range t.C {
		checkWifi()
	}
}

func checkWifi() {
	n := wifiname.WifiName()
	fmt.Println(n)
}
