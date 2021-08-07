module netease.com/dal

replace netease.com/conf => ../conf

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.0
	github.com/sirupsen/logrus v1.4.2
	go.mongodb.org/mongo-driver v1.7.0
)
