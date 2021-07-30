package domain

import (
	"context"
	"time"

	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"

	"netease.com/dal"
	"netease.com/dto"
)

var ctx = context.Background()
var json = jsoniter.ConfigCompatibleWithStandardLibrary
var UserLoginDomain = UserLoginDomainWrap{}

type UserLoginDomainWrap struct {
}

func (userLoginDomain *UserLoginDomainWrap) GetUserSession(account string) *dto.UserSession {
	redisKey := "Acc:" + account
	ret := dal.RedisCluster.Get(ctx, redisKey)
	log.Info("get from redis: ", ret.Val())

	userInfo := dto.UserSession{}
	json.Unmarshal([]byte(ret.Val()), &userInfo)
	return &userInfo
}

func (userLoginDomain *UserLoginDomainWrap) AddUserSession(userInfo *dto.UserInfo) *dto.UserSession {
	redisKey := "Acc:" + userInfo.Account

	userSession := dto.UserSession{}
	userSession.Account = userInfo.Account
	userSession.LoginTime = 1
	userSession.Sig = "mySig"

	ret, err1 := json.MarshalToString(&userSession)
	if err1 != nil {
		panic(err1)
	}
	err := dal.RedisCluster.Set(ctx, redisKey, ret, 5000*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return &userSession
}
