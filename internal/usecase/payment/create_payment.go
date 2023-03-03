package payment

import (
	"context"
	"errors"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var ()

func (p paymentUsecase) CreatePayment(ctx context.Context, arg pModel.AddPayment, au uModel.AuthUser) (pModel.Payment, error) {
	uuid, err := p.uuidGenerator.GenerateUUID()
	if err != nil {
		return pModel.Payment{}, err
	}

	switch arg.Method {
	case pModel.CashMethod:
		_, err = p.pRepository.CreatePayment(ctx, pModel.AddPayment{
			ID:      uuid,
			OrderID: "",
			Amount:  arg.Amount,
			Status:  pModel.UnpaidStatus,
			Method:  pModel.CashMethod,
		})
		if err != nil {
			return pModel.Payment{}, err
		}

		return p.pRepository.GetPayment(ctx, uuid)
	case pModel.QRISMethod:
		qris, err := p.xenditConnector.GenerateQRIS(uuid, arg.Amount)
		if err != nil {
			return pModel.Payment{}, err
		}

		id, err := p.pRepository.CreatePayment(ctx, pModel.AddPayment{
			ID:      qris.ExternalID,
			OrderID: "",
			Amount:  qris.Amount,
			Status:  pModel.UnpaidStatus,
			Method:  pModel.QRISMethod,
			QrStr:   qris.QRString,
		})
		if err != nil {
			return pModel.Payment{}, err
		}

		return p.pRepository.GetPayment(ctx, id)
	}

	return pModel.Payment{}, errors.New("")
}
