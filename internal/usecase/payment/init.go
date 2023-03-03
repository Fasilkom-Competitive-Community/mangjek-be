package payment

import (
	pRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/payment"
)

type paymentUsecase struct {
	pRepository     pRepo.Repository
	uuidGenerator   Generator
	xenditConnector PaymentConnector
}

func NewPaymentUsecase(pr pRepo.Repository, uG Generator, xC PaymentConnector) *paymentUsecase {
	return &paymentUsecase{pRepository: pr, uuidGenerator: uG, xenditConnector: xC}
}
