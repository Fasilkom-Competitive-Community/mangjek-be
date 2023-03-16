package qris

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
)

type Repository interface {
	CreateQRIS(ctx context.Context, arg pModel.AddQRIS) (pModel.QRIS, error)
	GetQRIS(ctx context.Context, externalID string) (pModel.QRIS, error)
}
