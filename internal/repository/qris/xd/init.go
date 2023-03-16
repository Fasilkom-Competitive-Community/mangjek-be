package xd

import (
	"context"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	qRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/qris"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/qrcode"
)

type xenditQRISRepository struct {
}

// CreateQRIS implements qris.QRIS
func (r xenditQRISRepository) CreateQRIS(ctx context.Context, arg pModel.AddQRIS) (pModel.QRIS, error) {
	req := qrcode.CreateQRCodeParams{
		ExternalID:  arg.ExternalID,
		Amount:      arg.Amount,
		Type:        xendit.DynamicQRCode,
		CallbackURL: "https://httpdump.app/inspect/15ef4794-3057-47e7-bb78-51778a5058a4",
	}
	// http://localhost:4001/payments/qris/callback

	resp, err := qrcode.CreateQRCodeWithContext(ctx, &req)
	if err != nil {
		return pModel.QRIS{}, err
	}

	return pModel.QRIS{
		ID:         resp.ID,
		ExternalID: resp.ExternalID,
		Amount:     resp.Amount,
		QrString:   resp.QRString,
		Status:     resp.Status,
	}, nil
}

// GetQRIS implements qris.QRIS
func (r xenditQRISRepository) GetQRIS(ctx context.Context, externalID string) (pModel.QRIS, error) {
	resp, err := qrcode.GetQRCodeWithContext(ctx, &qrcode.GetQRCodeParams{
		ExternalID: externalID,
	})
	if err != nil {
		return pModel.QRIS{}, err
	}

	return pModel.QRIS{
		ID:         resp.ID,
		ExternalID: resp.ExternalID,
		Amount:     resp.Amount,
		QrString:   resp.QRString,
		Status:     resp.Status,
	}, nil
}

func NewXenditQRISRepository() qRepo.Repository {
	return xenditQRISRepository{}
}
