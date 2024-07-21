package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/alexandermatseev/go_auth/internal/client/db"
	"github.com/alexandermatseev/go_auth/internal/model"
	"github.com/alexandermatseev/go_auth/internal/repository"
	"github.com/alexandermatseev/go_auth/internal/repository/user/converter"
	modelRepo "github.com/alexandermatseev/go_auth/internal/repository/user/model"
)

const (
	tableName = "users"

	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	roleColumn            = "role"
	passwordColumn        = "password"
	confirmPasswordColumn = "confirm_password"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, info *model.CreateUser) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn, confirmPasswordColumn).
		Values(info.Name, info.Email, info.Role, info.Password, info.ConfirmPassword).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, roleColumn, emailColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name, &user.Role, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Update(ctx context.Context, id int64, upd *model.UserUpdate) error {
	builder := sq.Update(tableName).PlaceholderFormat(sq.Dollar)
	if upd.Role != nil {
		builder = builder.Set(roleColumn, *upd.Role)
	}
	if upd.Name != nil {
		builder = builder.Set(nameColumn, *upd.Name)
	}
	if upd.Email != nil {
		builder = builder.Set(emailColumn, *upd.Email)
	}

	builder = builder.Set(updatedAtColumn, time.Now()).Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).Where(sq.Eq{idColumn: id})
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
