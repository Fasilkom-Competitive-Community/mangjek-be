package payment

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
)

func (p paymentUsecase) UpdatePaymentStatusToPaid(ctx context.Context, id string) (string, error) {
	return p.pRepository.UpdatePaymentStatus(ctx, id, pModel.PaidStatus)
}
