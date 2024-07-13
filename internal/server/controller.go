package server

import (
	"context"
	auth "github.com/alexandermatseev/go_auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Routes struct {
	auth.UnimplementedUserAuthServer
}

func (r *Routes) Get(_ context.Context, req *auth.GetRequest) (*auth.GetResponse, error) {
	resp := auth.GetResponse{
		Id:        req.Id,
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      auth.Role(gofakeit.RandomInt([]int{1, 2})),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}
	return &resp, nil
}

func (r *Routes) Create(_ context.Context, _ *auth.CreateRequest) (*auth.CreateResponse, error) {
	return &auth.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (r *Routes) Update(_ context.Context, _ *auth.UpdateRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (r *Routes) Delete(_ context.Context, _ *auth.DeleteRequest) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}
