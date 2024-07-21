package repository

import (
	"context"
	"github.com/alexandermatseev/go_auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, info *model.CreateUser) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, upd *model.UserUpdate) error
	Delete(ctx context.Context, id int64) error
}
