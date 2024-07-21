package model

import (
	"database/sql"
	"github.com/alexandermatseev/go_auth/internal/types"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Role      types.Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUser struct {
	Name            string
	Email           string
	Role            types.Role
	Password        []byte
	ConfirmPassword []byte
}

type UserUpdate struct {
	Role  *types.Role
	Name  *string
	Email *string
}
