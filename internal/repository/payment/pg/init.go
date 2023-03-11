package pg

import (
	"context"
	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	pRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/payment"
	"github.com/jackc/pgx/v4"
)

type pgPaymentRepository struct {
	querier sqlc.Querier
}

// CreatePayment implements payment.Payment
func (r pgPaymentRepository) CreatePayment(ctx context.Context, arg pModel.AddPayment) (string, error) {
	id, err := r.querier.CreatePayment(ctx, sqlc.CreatePaymentParams{
		ID:     arg.ID,
		Amount: arg.Amount,
		Status: string(arg.Status),
		Method: string(arg.Method),
		QrStr:  arg.QrString,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Payment not found")
	}
	return id, err
}

// GetPayment implements payment.Payment
func (r pgPaymentRepository) GetPayment(ctx context.Context, id string) (pModel.Payment, error) {
	p, err := r.querier.GetPayment(ctx, id)
	if err == pgx.ErrNoRows {
		return pModel.Payment{}, errorCommon.NewNotFoundError("Payment not found")
	}

	return pModel.Payment{
		ID:        p.ID,
		Amount:    p.Amount,
		Status:    pModel.Status(p.Status),
		Method:    pModel.Method(p.Method),
		QrString:  p.QrStr,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}

// UpdatePaymentStatus implements payment.Payment
func (r pgPaymentRepository) UpdatePaymentStatus(ctx context.Context, id string, status pModel.Status) (string, error) {
	id, err := r.querier.UpdatePaymentStatusToPaid(ctx, sqlc.UpdatePaymentStatusToPaidParams{
		ID:     id,
		Status: string(status),
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Payment not found")
	}
	return id, nil
}

func NewPGPaymentRepository(querier sqlc.Querier) pRepo.Repository {
	return pgPaymentRepository{querier: querier}
}
