package service

import (
	log "github.com/sirupsen/logrus"
	"netease.com/domain"
	"netease.com/dto"
	"netease.com/reqs"
)

var LoginService = LoginServiceWrap{}

type LoginServiceWrap struct {
}

func (loginServer *LoginServiceWrap) Login(login *reqs.LoginReq) *dto.UserSession {
	log.Info("----  username =", login.User)
	return domain.UserLoginDomain.GetUserSession(login.Account)
}

func (loginServer *LoginServiceWrap) Add(userInfo *dto.UserInfo) *dto.UserSession {
	log.Info("------ add player ", userInfo.UserName)
	return domain.UserLoginDomain.AddUserSession(userInfo)

}
