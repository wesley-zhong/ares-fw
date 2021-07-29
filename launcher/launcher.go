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
	webc := &controller.UserLogin{}

	//init rpc
	coreInst := core.Core{}
	coreInst.Init()
	controller.Register(&coreInst)
	coreInst.RegisterController(webc)
	//init dal
	dal.InitRedis("192.168.94.138:6379", "")
	//start tcp server
	gameServer := new(tcpServer.TcpServer)
	gameServer.Core = &coreInst
	go gameServer.Init("tcp4://:9000")
	//start web server
	webServer := webServer.AresWebapp{}
	webServer.WebAppStart(&coreInst, "localhost:8080")
}
