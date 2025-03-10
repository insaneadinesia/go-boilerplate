package driver

import (
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/config"
)

func NewAsynqClient(cfg config.Config) *asynq.Client {
	redisAddr := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: cfg.RedisPassword,
	})

	return client
}
