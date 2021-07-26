package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"netease.com/controller"
	"netease.com/core"
	tcpServer "netease.com/tcpserver"
)

func main() {
	log.SetOutput(os.Stdout)
	log.Info("main start")

	coreInst := core.Core{}
	coreInst.Init()
	controller.Register(&coreInst)
	gameServer := new(tcpServer.TcpServer)
	gameServer.Core = &coreInst
	gameServer.Init("tcp4://:9000")
}
