package user

import (
	"context"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Repository interface {
	CreateUser(ctx context.Context, arg uModel.AddUser) (string, error)
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (uModel.User, error)
	GetUserHistory(ctx context.Context, id string) ([]oModel.Order, error)
	VerifyAvailableUser(ctx context.Context, id string) (bool, error)
	ListUsers(ctx context.Context) ([]uModel.User, error)
	UpdateUser(ctx context.Context, arg uModel.UpdateUser) (string, error)
}
