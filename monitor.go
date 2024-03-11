package main

import (
	"log"
	"time"

	wifiname "git.okki.hu/garrichs/wifi-name"
)

func startOfficeMonitor(pollInterval int) {
	t := time.NewTicker(time.Duration(pollInterval) * time.Second)
	for range t.C {
		if usingOfficeWifi() {
			err := store.Flag()
			if err != nil {
				log.Println(err)
			}
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
