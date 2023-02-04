package user

import (
	"context"

	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
)

type Repository interface {
	CreateDriver(ctx context.Context, arg dModel.AddDriver) (int32, error)
	DeleteDriver(ctx context.Context, id int32) error
	GetDriver(ctx context.Context, id int32) (dModel.Driver, error)
	GetDriverByUserID(ctx context.Context, uid string) (dModel.Driver, error)
	VerifyAvailableDriver(ctx context.Context, uid string) (bool, error)
	ListDrivers(ctx context.Context) ([]dModel.Driver, error)
	UpdateDriver(ctx context.Context, arg dModel.UpdateDriver) (int32, error)
}
