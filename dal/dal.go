package dal

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Redis *redis.Client
var MongoClient *mongo.Client

func InitRedis(addr string, password string) {
	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
}

func InitMongodb(addr string, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+addr))
	MongoClient = client
	if err != nil {
		log.Error("connect ", addr, " failed")
	}
}

func getCollection(dbName string, connName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(connName)
}