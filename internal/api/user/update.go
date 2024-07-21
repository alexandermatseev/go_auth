package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/converter"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (r *Routes) Update(ctx context.Context, req *auth.UpdateRequest) (*empty.Empty, error) {
	err := r.userService.Update(ctx, req.GetId(), converter.ToServiceFromUpdateProto(req))
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
