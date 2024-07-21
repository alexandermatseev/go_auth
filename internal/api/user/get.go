package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/converter"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
)

func (r *Routes) Get(ctx context.Context, req *auth.GetRequest) (*auth.GetResponse, error) {
	user, err := r.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return converter.ToCreateProtoFromService(user), nil
}
