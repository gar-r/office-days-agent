package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var conf Conf
var store Store

func main() {
	envconfig.MustProcess(EnvConfigPrefix, &conf)
	store = NewFileStore(conf.DataPath)
	go func() {
		log.Println("starting office monitor...")
		startOfficeMonitor(conf.PollIntervalSeconds)
	}()
	log.Println("starting api server...")
	startApiServer(conf.ApiServerPort)
}
