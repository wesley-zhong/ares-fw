package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"netease.com/controller"
	"netease.com/core"
	tcpServer "netease.com/tcpserver"
	webServer "netease.com/webserver"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Info("main start")
	webc := &controller.UserLoginService{}

	coreInst := core.Core{}
	coreInst.Init()
	controller.Register(&coreInst)
	coreInst.RegisterController(webc)

	gameServer := new(tcpServer.TcpServer)
	gameServer.Core = &coreInst
	go gameServer.Init("tcp4://:9000")
	webServer := webServer.AresWebapp{}
	webServer.WebAppStart(&coreInst, "localhost:8080")
}
