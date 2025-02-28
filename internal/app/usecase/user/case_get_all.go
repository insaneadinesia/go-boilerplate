package user

import (
	"context"

	"github.com/insaneadinesia/go-boilerplate/internal/app/entity"
	"github.com/insaneadinesia/go-boilerplate/internal/app/repository"
	"github.com/insaneadinesia/go-boilerplate/internal/pkg/pagination"
	"golang.org/x/sync/errgroup"
)

func (u *usecase) GetAll(ctx context.Context, req GetAllUserRequest) (resp GetAllUserResponse, err error) {
	if req.Page == 0 {
		req.Page = 1
	}

	if req.PerPage == 0 {
		req.PerPage = 20
	}

	filter := repository.UserFilter{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Limit:    req.PerPage,
		Offset:   (req.Page - 1) * req.PerPage,
	}

	g := new(errgroup.Group)

	g.Go(func() (err error) {
		users, err := u.userRepository.GetAll(ctx, filter)
		if err != nil {
			return
		}

		if len(users) == 0 {
			resp.Users = []UserResponse{}
		}

		// Mapping Response
		for _, user := range users {
			resp.Users = append(resp.Users, u.mappingUserResponse(user).UserResponse)
		}

		return
	})

	g.Go(func() (err error) {
		total, err := u.userRepository.CountTotal(ctx, filter)
		if err != nil {
			return
		}

		// Mapping Response
		resp.Pagination = pagination.GeneratePaginationResponse(req.PerPage, req.Page, total)
		return
	})

	if err = g.Wait(); err != nil {
		return
	}

	return
}

func (u *usecase) mappingUserResponse(user entity.User) (result UserDetailResponse) {
	result = UserDetailResponse{
		UserResponse: UserResponse{
			UUID:      user.UUID.String(),
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:03:02Z0700"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:03:02Z0700"),
		},
	}

	return
}
