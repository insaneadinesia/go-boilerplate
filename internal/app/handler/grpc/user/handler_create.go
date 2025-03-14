package user

import (
	context "context"
	"fmt"

	"github.com/insaneadinesia/go-boilerplate/internal/app/usecase/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Create(ctx context.Context, req *CreateUserRequest) (resp *CreateUpdateDeleteResponse, err error) {
	in := user.CreateUpdateUserRequest{
		Name:          req.GetName(),
		Username:      req.GetUsername(),
		Email:         req.GetEmail(),
		SubDistrictID: req.GetSubDistrictId(),
	}

	if err = h.validator.Struct(in); err != nil {
		err = status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid request: %v", err))
		return
	}

	err = h.userUsecase.Create(ctx, in)
	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("failed to create user: %v", err))
		return
	}

	resp = &CreateUpdateDeleteResponse{
		Message: "Request Successfully Processed",
	}

	return
}
