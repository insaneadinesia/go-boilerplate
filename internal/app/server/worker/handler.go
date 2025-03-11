package worker

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/handler/worker/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/worker"
	"github.com/insaneadinesia/gobang/logger"
)

func SetupHandler(container *container.Container) *asynq.ServeMux {
	userWorker := user.NewHandler().SetUserUsecase(container.UserUsecase).Validate()

	mux := asynq.NewServeMux()

	mux.Use(TracingMiddleware)
	mux.Use(SetLoggingMiddleware(container))

	mux.HandleFunc(constants.QUEUE_USER_CREATED, Execute(userWorker.InformUserCreated))

	return mux
}

// Execute will extract the original payload and execute the function
func Execute(f func(ctx context.Context, payload any) (err error)) asynq.HandlerFunc {
	return func(ctx context.Context, task *asynq.Task) (err error) {
		defer func() {
			if err != nil {
				logger.Log.Error(ctx, "Execute Error", err.Error())
			}

			logger.Log.TDR(ctx)
		}()

		var payload worker.TaskPayload
		if err = json.Unmarshal(task.Payload(), &payload); err != nil {
			return
		}

		err = f(ctx, payload.Data)

		return
	}
}
