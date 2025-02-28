package user

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
)

func (u *usecase) Create(ctx context.Context, req CreateUpdateUserRequest) (err error) {
	user := entity.User{
		Name:          req.Name,
		Username:      req.Username,
		Email:         req.Email,
		SubDistrictID: req.SubDistrictID,
	}

	err = u.userRepository.Create(ctx, &user)
	return
}
