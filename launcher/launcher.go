package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"netease.com/conf"
	"netease.com/controller"
	"netease.com/core"
	"netease.com/dal"
	tcpServer "netease.com/tcpserver"
	webServer "netease.com/webserver"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Info("main start")
	conf.InitConf()

	//every modul init
	core.Init()

	//init msgId call
	controller.Init()

	dal.InitRedisCluster()
	dal.InitMongodb()
	//start tcp server
	gameServer := new(tcpServer.TcpServer)
	go gameServer.Init("tcp4://:9000")
	//start web server
	webServer := webServer.AresWebapp{}
	webServer.WebAppStart("localhost:8080")
}
