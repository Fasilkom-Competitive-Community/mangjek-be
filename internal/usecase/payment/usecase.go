package payment

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

type Usecase interface {
	Writer
	Reader
}

type Writer interface {
	CreatePayment(ctx context.Context, arg pModel.AddPayment, au uModel.AuthUser) (pModel.Payment, error)
	UpdatePaymentStatusToPaid(ctx context.Context, id string) (string, error)
}

type Reader interface {
	GetPayment(ctx context.Context, id string, au uModel.AuthUser) (pModel.Payment, error)
}

type Generator interface {
	GenerateUUID() (string, error)
}
