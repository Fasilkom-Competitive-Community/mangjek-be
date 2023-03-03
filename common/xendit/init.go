package xendit

import (
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/client"
	"github.com/xendit/xendit-go/qrcode"
)

type XenditClient struct {
	client *client.API
}

func NewXenditClient(secretKey string) *XenditClient {
	return &XenditClient{client: client.New(secretKey)}
}

func (x *XenditClient) GenerateQRIS(externalID string, amount float64) (pModel.QRIS, error) {
	req := qrcode.CreateQRCodeParams{
		ExternalID:  externalID,
		Type:        xendit.DynamicQRCode,
		CallbackURL: "https://httpdump.app/inspect/15ef4794-3057-47e7-bb78-51778a5058a4",
		Amount:      amount,
	}
	// http://localhost:4001/payments/qris/callback

	resp, err := x.client.QRCode.CreateQRCode(&req)
	if err != nil {
		return pModel.QRIS{}, err
	}

	return pModel.QRIS{
		ID:         resp.ID,
		ExternalID: resp.ExternalID,
		Amount:     resp.Amount,
		QRString:   resp.QRString,
		Status:     resp.Status,
	}, nil
}

func (x *XenditClient) ReadQRIS(externalID string) (pModel.QRIS, error) {
	resp, err := x.client.QRCode.GetQRCode(&qrcode.GetQRCodeParams{
		ExternalID: externalID,
	})
	if err != nil {
		return pModel.QRIS{}, err
	}

	return pModel.QRIS{
		ID:         resp.ID,
		ExternalID: resp.ExternalID,
		Amount:     resp.Amount,
		QRString:   resp.QRString,
		Status:     resp.Status,
	}, nil
}
