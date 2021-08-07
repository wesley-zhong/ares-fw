package conf

import (
	"strings"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/sirupsen/logrus"
)

func ApolloInit() {
	apolloCnf := &Config.ApolloConf
	logrus.Info("apollo addr =", apolloCnf.IP)

	c := &config.AppConfig{
		AppID:          apolloCnf.AppID,
		Cluster:        apolloCnf.Cluster,
		IP:             apolloCnf.IP,
		NamespaceName:  apolloCnf.NamespaceName,
		IsBackupConfig: apolloCnf.IsBackupConfig,
		Secret:         apolloCnf.Secret,
	}
	agollo.SetLogger(logrus.New())

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("apollo init success")
	nameSpaces := strings.Split(apolloCnf.NamespaceName, ",")
	for _, n := range nameSpaces {
		logrus.Info("==================== start namespace = ", n, "  ======================")
		cache := client.GetConfigCache(n)
		//var packLimit, _ = cache.Get("packet.limit")
		//logrus.Info("-----packet.limit  -get from apollo =", packLimit)
		checkKey(cache, client)
	}

}

func checkKey(cache agcache.CacheInterface, client *agollo.Client) {
	count := 0
	cache.Range(func(key, value interface{}) bool {
		logrus.Info("key : ", key, ", value :", value)
		count++
		return true
	})
	if count < 1 {
		panic("config key can not be null")
	}
}
