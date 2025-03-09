package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/apperror"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
	"gorm.io/gorm"
)

func (u *usecase) Delete(ctx context.Context, reqUUID string) (err error) {
	ctx, span := gotel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	parseUUID, err := uuid.Parse(reqUUID)
	if err != nil {
		return
	}

	user, err := u.userRepository.GetByUUID(ctx, parseUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = apperror.New(http.StatusNotFound, constants.CODE_RECORD_NOT_FOUND, err)
		}

		return
	}

	err = u.userRepository.Delete(ctx, &user)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, constants.CODE_DELETE_ERROR, err)
		return
	}

	return
}
