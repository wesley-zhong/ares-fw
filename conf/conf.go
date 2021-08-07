package conf

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config ConfigWrap

func InitConf() {
	viper.SetConfigName("common")
	viper.AddConfigPath("../resources")
	err := viper.ReadInConfig()
	if err != err {
		log.Error("read config failed", err)
		panic(err)
	}
	viper.Unmarshal(&Config)
	log.Info("load redis cluster addr = ", Config.Rediscluster.Addr)
}

type ConfigWrap struct {
	Rediscluster RedisClusterConf
	Mongodb      MongoDBConf
}

type RedisClusterConf struct {
	Addr     string
	User     string
	Password string
}

type MongoDBConf struct {
	Addr     string
	User     string
	Password string
}

func LoadConf(conf interface{}) error {
	return nil
}

func LoadRedisConf() *RedisClusterConf {
	reddisClusterConf := &RedisClusterConf{}

	return reddisClusterConf
}

func LoadMongodConf() *MongoDBConf {
	mongdbConf := &MongoDBConf{}

	return mongdbConf
}
