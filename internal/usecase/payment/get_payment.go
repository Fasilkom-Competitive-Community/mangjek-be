package payment

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (p paymentUsecase) GetPayment(ctx context.Context, id string, au uModel.AuthUser) (pModel.Payment, error) {
	return p.pRepository.GetPayment(ctx, id)
}
