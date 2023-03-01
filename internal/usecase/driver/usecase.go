package user

import (
	"context"

	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Usecase interface {
	CreateDriver(ctx context.Context, arg dModel.AddDriver, au uModel.AuthUser) (int32, error)
	DeleteDriver(ctx context.Context, id int32, uid string, au uModel.AuthUser) error
	GetDriver(ctx context.Context, id int32, au uModel.AuthUser) (dModel.Driver, error)
	ListDrivers(ctx context.Context, au uModel.AuthUser) ([]dModel.Driver, error)
	UpdateDriver(ctx context.Context, arg dModel.UpdateDriver, au uModel.AuthUser) (int32, error)
}
