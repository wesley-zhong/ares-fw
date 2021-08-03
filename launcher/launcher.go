package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"netease.com/controller"
	"netease.com/core"
	"netease.com/dal"
	tcpServer "netease.com/tcpserver"
	webServer "netease.com/webserver"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Info("main start")

	//every modul init
	core.Init()

	//init msgId call
	controller.Init()

	dal.InitRedisCluster("192.168.94.138:6379", "")
	dal.InitMongodb("192.168.94.138:27017", "dice_dev_user", "3cdY61jr")
	//start tcp server
	gameServer := new(tcpServer.TcpServer)
	go gameServer.Init("tcp4://:9000")
	//start web server
	webServer := webServer.AresWebapp{}
	webServer.WebAppStart("localhost:8080")
}
