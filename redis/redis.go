package redis

import (
	"github.com/go-redis/redis"
	"misteraladin.com/jasmine/go-boiler-plate/config"
)

var (
	redisConfig = config.Config.Redis
	redisClient *redis.Client
)

func init() {
	setupRedisConn()
}

func setupRedisConn() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})
}

func RedisClient() *redis.Client {
	return redisClient
}
