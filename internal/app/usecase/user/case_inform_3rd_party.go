package user

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
	"github.com/insaneadinesia/gobang/logger"
)

func (u *usecase) Inform3rdParty(ctx context.Context, req CreatedUserPayload) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	// DO SOMETHING
	logger.Log.Info(ctx, "Inform 3rd Party...", req)

	return
}
