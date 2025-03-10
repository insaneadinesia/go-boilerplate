package worker

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/app/handler/worker/user"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
)

func SetupHandler(container *container.Container) asynq.HandlerFunc {
	userWorker := user.NewHandler().SetUserUsecase(container.UserUsecase).Validate()

	return func(ctx context.Context, t *asynq.Task) (err error) {
		switch t.Type() {
		case constants.QUEUE_USER_CREATED:
			err = userWorker.InformUserCreated(ctx, t)
		default:
			err = fmt.Errorf("unexpected task type: %s", t.Type())
		}

		return
	}
}
