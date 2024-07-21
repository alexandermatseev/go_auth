package user

import (
	"github.com/alexandermatseev/go_auth/internal/client/db"
	"github.com/alexandermatseev/go_auth/internal/repository"
	"github.com/alexandermatseev/go_auth/internal/service"
)

type userService struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewUserService(userRepository repository.UserRepository, txManager db.TxManager) service.UserService {
	return &userService{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
