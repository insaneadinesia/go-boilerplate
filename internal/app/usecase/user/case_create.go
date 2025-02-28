package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/apperror"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/constants"
	"gorm.io/gorm"
)

func (u *usecase) Create(ctx context.Context, req CreateUpdateUserRequest) (err error) {
	user := entity.User{
		Name:          req.Name,
		Username:      req.Username,
		Email:         req.Email,
		SubDistrictID: req.SubDistrictID,
	}

	existingUser, err := u.userRepository.GetByUsernameOrEmail(ctx, req.Username, req.Email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if existingUser.UUID.String() != "" {
		err = apperror.New(http.StatusUnprocessableEntity, constants.CODE_DUPLICATE_USERNAME_OR_EMAIL, errors.New("duplicate username or email"))

		return
	}

	err = u.userRepository.Create(ctx, &user)
	return
}
