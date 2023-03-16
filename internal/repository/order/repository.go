package order

import (
	"context"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
)

type Repository interface {
	CreateOrderInquiry(ctx context.Context, arg oModel.AddOrderInquiry) (string, error)
	GetOrderInquiry(ctx context.Context, id string) (oModel.OrderInquiry, error)
	DeleteOrderInquiry(ctx context.Context, id string) error

	CreateOrder(ctx context.Context, payment pModel.AddPayment, order oModel.AddOrder) (string, error)
	GetOrder(ctx context.Context, id string) (oModel.Order, error)
}
