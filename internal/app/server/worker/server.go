package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
)

func StartWorkerService(container *container.Container) {
	cfg := container.Config

	server := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
			Password: cfg.RedisPassword,
		},
		asynq.Config{
			Concurrency: 10,
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				retried, _ := asynq.GetRetryCount(ctx)
				maxRetry, _ := asynq.GetMaxRetry(ctx)

				if retried >= maxRetry {
					// DO SOMETHING WHEN MAX RETRY REACHED
					return
				}
			}),
		},
	)

	fmt.Println("Worker service is running ...")

	if err := server.Run(SetupHandler(container)); err != nil {
		panic(err)
	}
}
