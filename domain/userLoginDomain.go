package domain

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

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
	val, err := dal.RedisCluster.Get(ctx, redisKey).Result()
	if err == redis.Nil || err != nil {
		log.Info("redis key =", redisKey, "  not exist")
		//get from mongodb
		dal.GetCollection("account", "account").FindOne(context.Background(), bson.M{"account": account})
	}
	log.Info("get from redis: ", val)
	userInfo := dto.UserSession{}
	json.Unmarshal([]byte(val), &userInfo)
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
	var upset = true
	retM, errM := dal.GetCollection("account", "account").ReplaceOne(context.Background(), bson.M{"account": userInfo.Account}, userSession, &options.ReplaceOptions{Upsert: &upset})
	if errM == nil {
		log.Info("mongdb update success matched count =", retM.MatchedCount)
	}

	err := dal.RedisCluster.Set(ctx, redisKey, ret, 5000*time.Second).Err()
	if err != nil {
		panic(err)
	}
	return &userSession
}
