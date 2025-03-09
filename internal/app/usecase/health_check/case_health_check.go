package health_check

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
)

func (u *usecase) HealthCheck(ctx context.Context) (resp StatusCheck, err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	const (
		STATUS_OK    = "OK"
		STATUS_ERROR = "ERROR"
	)

	resp.DBStatus = STATUS_OK

	err = u.healthCheckRepository.PingDB(ctx)
	if err != nil {
		resp.DBStatus = STATUS_ERROR
		return
	}

	return
}
