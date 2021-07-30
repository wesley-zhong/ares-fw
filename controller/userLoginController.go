package controller

import (
	"fmt"

	"netease.com/dto"
	"netease.com/reqs"
	"netease.com/service"
)

/**
note: this is resultful call such as  ip:port/UserLogin/Login   and  with json body   reqs.Login
**/

type UserLogin struct {
}

func (*UserLogin) Login(login *reqs.LoginReq) interface{} {
	return service.LoginService.Login(login)
}

func (*UserLogin) Update(userInfo *dto.UserInfo) interface{} {
	fmt.Println("user info update username =" + userInfo.UserName)
	return service.LoginService.Add(userInfo)
}
