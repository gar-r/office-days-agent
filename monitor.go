package main

import (
	"log"
	"time"

	wifiname "github.com/gar-r/wifi-name"
)

const DateFormat = "20060102"

func startOfficeMonitor(pollInterval int) {
	t := time.NewTicker(time.Duration(pollInterval) * time.Second)
	for range t.C {
		date := time.Now().Format(DateFormat)
		od, err := store.Lookup(date)
		if err != nil {
			log.Println(err)
			continue
		}
		if od {
			continue // already logged office
		}
		err = store.Save(date, usingOfficeWifi())
		if err != nil {
			log.Println(err)
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
