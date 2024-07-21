package user

import (
	"github.com/alexandermatseev/go_auth/internal/service"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
)

type Routes struct {
	auth.UnimplementedUserAuthServer
	userService service.UserService
}

func NewRoutes(userService service.UserService) *Routes {
	return &Routes{
		userService: userService,
	}
}
