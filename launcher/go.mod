module netease.com/launcher

replace netease.com/core => ../core

replace netease.com/controller => ../controller

replace netease.com/gen => ../gen

replace netease.com/tcpserver => ../tcpserver

go 1.16

require (
	github.com/panjf2000/gnet v1.4.5 // indirect
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	netease.com/controller v0.0.0-00010101000000-000000000000
	netease.com/core v0.0.0-00010101000000-0000000000001
	netease.com/gen v0.0.0-00010101000000-000000000000 // indirect
	netease.com/tcpserver v0.0.0-00010101000000-000000000000

)
