package dal

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"netease.com/conf"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Redis *redis.Client
var MongoClient *mongo.Client
var RedisCluster *redis.ClusterClient

func InitRedisCluster() {
	redisCluserConf := &conf.Config.Rediscluster
	initRedisCluster(redisCluserConf.Addr, redisCluserConf.Password)
}

func InitMongodb() {
	mongoDBConf := &conf.Config.Mongodb
	initMongodb(mongoDBConf.Addr, mongoDBConf.User, mongoDBConf.Password)

}

func InitRedis(addr string, password string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func initRedisCluster(addr string, password string) {
	RedisCluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{addr},

		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
	})
}

func initMongodb(addr string, account string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Client()
	opts.ApplyURI("mongodb://" + addr)
	opts.Auth = &options.Credential{}
	opts.Auth.Username = account
	opts.Auth.Password = password
	client, err := mongo.Connect(ctx, opts)
	MongoClient = client
	if err != nil {
		log.Error("connect ", addr, " failed")
	}
}

func GetCollection(dbName string, connName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(connName)
}
