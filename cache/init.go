package cache

import (
	"github.com/go-redis/redis"
	"theychat/config"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr: config.Conf.RedisConfig.Host + ":" + config.Conf.RedisConfig.Port,
		DB:   config.Conf.RedisConfig.DB,
	})
	_, err := RDB.Ping().Result()
	if err != nil {
		panic(err)
	}
}
