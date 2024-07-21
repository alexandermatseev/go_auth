package user

import (
	"context"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (r *Routes) Delete(ctx context.Context, req *auth.DeleteRequest) (*empty.Empty, error) {
	err := r.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
