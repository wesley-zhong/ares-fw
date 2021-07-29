module netease.com/launcher

replace netease.com/core => ../core

replace netease.com/controller => ../controller

replace netease.com/gen => ../gen

replace netease.com/tcpserver => ../tcpserver

replace netease.com/webserver => ../webserver

replace netease.com/reqs => ../reqs

replace netease.com/dal => ../dal

replace netease.com/dto => ../dto

replace netease.com/service => ../service

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	netease.com/controller v0.0.0-00010101000000-000000000000
	netease.com/core v0.0.0-00010101000000-0000000000001
	netease.com/dal v0.0.0-00010101000000-000000000000
	netease.com/tcpserver v0.0.0-00010101000000-000000000000
	netease.com/webserver v0.0.0-00010101000000-000000000000

)
