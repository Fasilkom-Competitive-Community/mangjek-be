package user

import (
	"context"

	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Usecase interface {
	CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error)
	DeleteUser(ctx context.Context, id string, au uModel.AuthUser) error
	GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error)
	ListUsers(ctx context.Context, au uModel.AuthUser) ([]uModel.User, error)
	UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error)
}
