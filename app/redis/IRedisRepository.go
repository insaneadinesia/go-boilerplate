package redis

type IRedisRepository interface {
	Set(key string, value interface{}) error
	Delete(keys []string) error
}
