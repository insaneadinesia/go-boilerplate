package user

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/app/repository"
)

type UserUsecase interface {
	Create(ctx context.Context, req CreateUpdateUserRequest) (err error)
	GetAll(ctx context.Context, req GetAllUserRequest) (resp GetAllUserResponse, err error)
}

type usecase struct {
	userRepository repository.User
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) SetUserRepository(repo repository.User) *usecase {
	u.userRepository = repo
	return u
}

func (u *usecase) Validate() UserUsecase {
	if u.userRepository == nil {
		panic("userRepository is nil")
	}

	return u
}
