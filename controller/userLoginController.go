package controller

import (
	log "github.com/sirupsen/logrus"
	"netease.com/dto"
	"netease.com/reqs"
	"netease.com/service"
)

/*
note: this is resultful call such as  ip:port/UserLogin/Login   and  with json body   reqs.Log
**/

type UserLogin struct {
}

func (userLogin *UserLogin) Login(login *reqs.LoginReq) interface{} {
	return service.LoginService.Login(login)
}

func (userLogin *UserLogin) Update(userInfo *dto.UserInfo) interface{} {
	log.Info("user info update username =", userInfo.UserName)
	return service.LoginService.Add(userInfo)
}
