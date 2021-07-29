package service

import (
	log "github.com/sirupsen/logrus"
	"netease.com/dto"
	"netease.com/reqs"
)

var LoginService = loginService{}

type loginService struct {
}

func (longServer *loginService) Login(login *reqs.LoginReq) *dto.UserInfo {
	log.Info("----  username =", login.User)
	userInfo := dto.UserInfo{}
	userInfo.NickName = "hello:" + login.User
	return &userInfo
}
