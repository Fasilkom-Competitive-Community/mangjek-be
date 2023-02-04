package driver

import (
	"context"

	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
)

type Repository interface {
	CreateDriver(ctx context.Context, arg uModel.AddDriver) (string, error)
	Delete(ctx context.Context, id string) error
	GetDriver(ctx context.Context, id string) (uModel.Driver, error)
	VerifyAvailableDriver(ctx context.Context, id string) (bool, error)
	ListDrivers(ctx context.Context) ([]uModel.Driver, error)
	UpdateDriver(ctx context.Context, arg uModel.UpdateDriver) (string, error)
}
