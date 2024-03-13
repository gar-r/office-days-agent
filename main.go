package main

import (
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
)

var conf Conf
var store DB

func main() {
	setup()
	go func() {
		log.Println("starting office monitor...")
		startOfficeMonitor(conf.PollIntervalSeconds)
	}()
	log.Println("starting api server...")
	startApiServer(conf.ApiServerPort)
}

func setup() {
	envconfig.MustProcess(EnvConfigPrefix, &conf)
	conf.DBPath = os.ExpandEnv(conf.DBPath)
	store = NewDiskDB(conf.DBPath)
}
