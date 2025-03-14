package user

import (
	context "context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *handler) GetDetail(ctx context.Context, req *GetUserDetailRequest) (resp *GetUserDetailResponse, err error) {
	uuid := req.GetUuid()
	if err = h.validator.Var(uuid, "required,uuid"); err != nil {
		err = status.Errorf(codes.InvalidArgument, fmt.Sprintf("invalid request: %v", err))
		return
	}

	out, err := h.userUsecase.GetDetail(ctx, uuid)
	if err != nil {
		err = status.Errorf(codes.Internal, fmt.Sprintf("failed to get user detail: %v", err))
		return
	}

	resp = &GetUserDetailResponse{
		Message: "Request Successfully Processed",
		Data: &UserData{
			Uuid:      out.UUID,
			Name:      out.Name,
			Username:  out.Username,
			Email:     out.Email,
			CreatedAt: out.CreatedAt,
			UpdatedAt: out.UpdatedAt,
		},
	}

	return
}
