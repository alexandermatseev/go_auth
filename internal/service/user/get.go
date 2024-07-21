package user

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/model"
)

func (s *userService) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
