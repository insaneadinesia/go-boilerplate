package user

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/insaneadinesia/go-boilerplate/internal/app/repository"
	"github.com/insaneadinesia/go-boilerplate/internal/app/wrapper/location_svc"
)

type UserUsecase interface {
	Create(ctx context.Context, req CreateUpdateUserRequest) (err error)
	GetAll(ctx context.Context, req GetAllUserRequest) (resp GetAllUserResponse, err error)
	GetDetail(ctx context.Context, reqUUID string) (resp UserDetailResponse, err error)
	Update(ctx context.Context, reqUUID string, req CreateUpdateUserRequest) (err error)
	Delete(ctx context.Context, reqUUID string) (err error)
	Inform3rdParty(ctx context.Context, req CreatedUserPayload) (err error)
}

type usecase struct {
	userRepository     repository.User
	locationSvcWrapper location_svc.LocationScvWrapper
	asynqClient        *asynq.Client
}

func NewUsecase() *usecase {
	return &usecase{}
}

func (u *usecase) SetUserRepository(repo repository.User) *usecase {
	u.userRepository = repo
	return u
}

func (u *usecase) SetLocationSvcWrapper(wrapper location_svc.LocationScvWrapper) *usecase {
	u.locationSvcWrapper = wrapper
	return u
}

func (u *usecase) SetAsynqClient(client *asynq.Client) *usecase {
	u.asynqClient = client
	return u
}

func (u *usecase) Validate() UserUsecase {
	if u.userRepository == nil {
		panic("userRepository is nil")
	}

	if u.locationSvcWrapper == nil {
		panic("locationSvcWrapper is nil")
	}

	if u.asynqClient == nil {
		panic("asynqClient is nil")
	}

	return u
}
