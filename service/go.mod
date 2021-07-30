module netease.com/service

replace netease.com/reqs => ../reqs

replace netease.com/dto => ../dto

replace netease.com/domain => ../domain

go 1.16

require (
	github.com/sirupsen/logrus v1.8.1
	netease.com/dto v0.0.0-00010101000000-000000000000
	netease.com/reqs v0.0.0-00010101000000-000000000000
)
