package user

import (
	context "context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) Delete(ctx context.Context, req *DeleteUserRequest) (resp *CreateUpdateDeleteResponse, err error) {
	uuid := req.GetUuid()
	if err = h.validator.Var(uuid, "required,uuid"); err != nil {
		err = status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid request: %v", err))
		return
	}

	err = h.userUsecase.Delete(ctx, uuid)
	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("failed to delete user: %v", err))
		return
	}

	resp = &CreateUpdateDeleteResponse{
		Message: "Request Successfully Processed",
	}

	return
}
