package converter

import (
	"github.com/alexandermatseev/go_auth/internal/model"
	"github.com/alexandermatseev/go_auth/internal/types"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateProtoFromService(user *model.User) *auth.GetResponse {

	return &auth.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      auth.Role(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt.Time),
	}
}
func ToServiceFromUpdateProto(user *auth.UpdateRequest) *model.UserUpdate {
	role := types.Role(user.Role)
	return &model.UserUpdate{
		Name:  &user.Name.Value,
		Email: &user.Email.Value,
		Role:  &role,
	}
}

func ToUserFromProto(name, email, password, confirmPassword string, role int) (*model.CreateUser, error) {
	user := &model.CreateUser{
		Name:            name,
		Email:           email,
		Role:            types.Role(role),
		Password:        password,
		ConfirmPassword: confirmPassword,
	}
	err := validateCreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func validateCreateUser(user *model.CreateUser) error {
	if user.Name == "" {
		return errors.New("Name is required")
	}
	if user.Email == "" {
		return errors.New("Email is required")
	}
	if user.Password == "" {
		return errors.New("Password is required")
	}
	if user.ConfirmPassword == "" {
		return errors.New("ConfirmPassword is required")
	}
	return nil
}
