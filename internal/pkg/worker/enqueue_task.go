package worker

import (
	"context"
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/gobang/gotel"
	"go.opentelemetry.io/otel/propagation"
)

type TaskPayload struct {
	Tracing map[string]string
	Data    any
}

func AsynqEnqueueTaskWithTrace(ctx context.Context, client *asynq.Client, taskName string, data any, maxRetry int) (err error) {
	if client == nil {
		return
	}

	// Inject the current trace context into a carrier
	carrier := make(propagation.MapCarrier)
	propagator := gotel.GetTextMapPropagator()
	propagator.Inject(ctx, carrier)

	// Create payload with data and tracing info
	payload := TaskPayload{
		Data:    data,
		Tracing: carrier,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Enqueue the task
	task := asynq.NewTask(taskName, payloadBytes, asynq.MaxRetry(maxRetry))
	_, err = client.EnqueueContext(ctx, task)

	return
}
