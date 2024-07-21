package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/model"
)

func (s *userService) Update(ctx context.Context, id int64, user *model.UserUpdate) error {
	err := s.userRepository.Update(ctx, id, user)
	if err != nil {
		return err
	}
	return nil
}
