package payment

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
)

type Repository interface {
	CreatePayment(ctx context.Context, arg pModel.AddPayment) (string, error)
	GetPayment(ctx context.Context, id string) (pModel.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id string, status pModel.Status) (string, error)
}
