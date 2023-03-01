package order

import (
	"context"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Usecase interface {
	Writer
	Reader
}

type Writer interface {
	CreateOrderInquiry(ctx context.Context, arg oModel.AddOrderInquiry, au uModel.AuthUser) (oModel.Direction, oModel.OrderInquiry, error)
	DeleteOrderInquiry(ctx context.Context, id string, au uModel.AuthUser) error
}

type Reader interface {
	GetOrderInquiry(ctx context.Context, id string, au uModel.AuthUser) (oModel.OrderInquiry, error)
}

type Calculator interface {
	CalculateDirection(ctx context.Context, origin oModel.Location, destination oModel.Location) (oModel.Direction, error)
}

type Generator interface {
	GenerateUUID() (string, error)
}
