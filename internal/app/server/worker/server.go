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
				fmt.Printf("task failed, error: %v\n", err)

				retried, _ := asynq.GetRetryCount(ctx)
				maxRetry, _ := asynq.GetMaxRetry(ctx)

				// You can put some dead letter queue here, like sending notification to internal team
				if retried >= maxRetry {
					fmt.Println("retry exhausted")
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
