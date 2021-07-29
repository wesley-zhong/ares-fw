module netease.com/controller

replace netease.com/core => ../core

replace netease.com/gen => ../gen

replace netease.com/reqs => ../reqs

replace netease.com/service => ../service

replace netease.com/dto => ../dto

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/protobuf v1.27.1
	netease.com/core v0.0.0-00010101000000-000000000000
	netease.com/dto v0.0.0-00010101000000-000000000000
	netease.com/gen v0.0.0-00010101000000-000000000000
	netease.com/reqs v0.0.0-00010101000000-000000000000
	netease.com/service v0.0.0-00010101000000-000000000000
)
