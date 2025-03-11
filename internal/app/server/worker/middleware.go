package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/container"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/worker"
	"github.com/insaneadinesia/gobang/gotel"
	"github.com/insaneadinesia/gobang/logger"
	"go.opentelemetry.io/otel/propagation"
)

// TracingMiddleware will extract the tracing context from the payload and start a span
func TracingMiddleware(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		var payload worker.TaskPayload

		if err := json.Unmarshal(task.Payload(), &payload); err != nil {
			// Log error but proceed without tracing
			return next.ProcessTask(ctx, task)
		}

		// Extract trace context
		propagator := gotel.GetTextMapPropagator()
		ctx = propagator.Extract(ctx, propagation.MapCarrier(payload.Tracing))

		// Start span
		tracer := gotel.DefaultTracerProvider().Tracer("asynq-worker")
		spanName := fmt.Sprintf("WORKER %s", task.Type())
		ctx, span := tracer.Start(ctx, spanName)
		defer span.End()

		// Proceed with the task handling
		return next.ProcessTask(ctx, task)
	})
}

func SetLoggingMiddleware(container *container.Container) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			cfg := container.Config
			ctxLogger := logger.Context{
				ServiceName:    cfg.AppName,
				ServiceVersion: cfg.AppVersion,
				ServicePort:    cfg.AppHTTPPort,
				ReqMethod:      "WORKER",
				ReqURI:         task.Type(),
				ReqBody:        task.Payload(),
			}

			// Inject Logger Context To Original Context
			ctx = logger.InjectCtx(ctx, ctxLogger)

			return next.ProcessTask(ctx, task)
		})
	}
}
