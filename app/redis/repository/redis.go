package repository

import (
	"time"

	"misteraladin.com/jasmine/go-boiler-plate/config"

	"github.com/go-redis/redis"
	RedisInterface "misteraladin.com/jasmine/go-boiler-plate/app/redis"
)

var redisDuration = config.Config.Redis.Duration

type RedisRepository struct {
	Client *redis.Client
}

func NewRedisRepository(Client *redis.Client) RedisInterface.IRedisRepository {
	return &RedisRepository{Client}
}

func (m *RedisRepository) Get(key string) (string, error) {
	val, err := m.Client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return val, nil
}

func (m *RedisRepository) Set(key string, value interface{}) error {
	err := m.Client.Set(key, value, time.Duration(redisDuration)*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

func (m *RedisRepository) Delete(keys []string) error {
	err := m.Client.Del(keys...).Err()
	if err != nil {
		return err
	}

	return nil
}
