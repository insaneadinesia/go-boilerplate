package user

import (
	context "context"
	"fmt"

	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetAll(ctx context.Context, req *GetAllUserRequest) (resp *GetAllUserResponse, err error) {
	in := user.GetAllUserRequest{
		Name:     req.GetName(),
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Page:     int(req.GetPage()),
		PerPage:  int(req.GetPerPage()),
	}

	out, err := h.userUsecase.GetAll(ctx, in)
	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("failed to get all user: %v", err))
		return
	}

	resp = &GetAllUserResponse{
		Message: "Request Successfully Processed",
		Data: &GetAllUserData{
			Users: []*UserData{},
			Pagination: &PaginationData{
				Page:       int32(out.Pagination.Page),
				PerPage:    int32(out.Pagination.PerPage),
				PageCount:  int32(out.Pagination.PageCount),
				TotalCount: out.Pagination.TotalCount,
			},
		},
	}

	for _, user := range out.Users {
		resp.Data.Users = append(resp.Data.Users, &UserData{
			Uuid:      user.UUID,
			Name:      user.Name,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return
}
