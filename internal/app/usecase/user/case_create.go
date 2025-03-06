package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/apperror"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/helper"
	"github.com/insaneadinesia/gobang/gotel"
	"gorm.io/gorm"
)

func (u *usecase) Create(ctx context.Context, req CreateUpdateUserRequest) (err error) {
	ctx, span := gotel.Otel.DefaultTracer().Start(ctx, helper.GetFuncName())
	defer span.End()

	user := entity.User{
		Name:          req.Name,
		Username:      req.Username,
		Email:         req.Email,
		SubDistrictID: req.SubDistrictID,
	}

	existingUser, err := u.userRepository.GetByUsernameOrEmail(ctx, req.Username, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	if existingUser.Username != "" {
		err = apperror.New(http.StatusUnprocessableEntity, constants.CODE_DUPLICATE_USERNAME_OR_EMAIL, errors.New("duplicate username or email"))

		return
	}

	err = u.userRepository.Create(ctx, &user)
	if err != nil {
		err = apperror.New(http.StatusUnprocessableEntity, constants.CODE_CREATE_ERROR, err)
		return
	}

	return
}
