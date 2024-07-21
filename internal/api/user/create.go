package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/converter"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
)

func (r *Routes) Create(ctx context.Context, req *auth.CreateRequest) (*auth.CreateResponse, error) {
	createUser, err := converter.ToUserFromProto(req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		int(req.GetRole()))
	if err != nil {
		return nil, err
	}
	id, err := r.userService.Create(ctx, createUser)

	if err != nil {
		return nil, err
	}
	return &auth.CreateResponse{Id: id}, nil
}
