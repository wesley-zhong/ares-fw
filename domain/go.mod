module netease.com/domain

replace netease.com/dal => ../dal

replace netease.com/dto => ../dto

go 1.16

require (
	github.com/json-iterator/go v1.1.11
	github.com/sirupsen/logrus v1.8.1
	netease.com/dal v0.0.0-00010101000000-000000000000
	netease.com/dto v0.0.0-00010101000000-000000000000
)
