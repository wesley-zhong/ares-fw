module netease.com/tcpserver

replace netease.com/core => ../core

replace netease.com/gen => ../gen

go 1.16

require (
	github.com/golang/protobuf v1.5.2
	github.com/panjf2000/gnet v1.5.3
	github.com/sirupsen/logrus v1.8.1
	go.uber.org/atomic v1.9.0 // indirect
	netease.com/core v0.0.0-00010101000000-000000000000
)
