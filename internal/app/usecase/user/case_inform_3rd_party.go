package user

import (
	"context"
	"fmt"

	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
)

func (u *usecase) Inform3rdParty(ctx context.Context, req CreatedUserPayload) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	// DO SOMETHING
	fmt.Println("Created User UUID: ", req.UUID)
	fmt.Println("Inform 3rd Party...")

	return
}
