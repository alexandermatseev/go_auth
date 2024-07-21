package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/model"
	"github.com/alexandermatseev/go_auth/internal/utils"
	"github.com/pkg/errors"
)

func (s *userService) Create(ctx context.Context, user *model.CreateUser) (int64, error) {
	isCompare := utils.ComparePasswords(user.Password, user.ConfirmPassword)
	if !isCompare {
		return 0, errors.New("Password is not equal")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword
	confirmHashedPassword, err := utils.HashPassword(user.ConfirmPassword)
	if err != nil {
		return 0, err
	}
	user.ConfirmPassword = confirmHashedPassword
	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
