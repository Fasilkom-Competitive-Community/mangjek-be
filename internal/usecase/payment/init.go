package payment

import (
	pRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/payment"
	qRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/qris"
)

type paymentUsecase struct {
	pRepository   pRepo.Repository
	qRepository   qRepo.Repository
	uuidGenerator Generator
}

func NewPaymentUsecase(pr pRepo.Repository, qr qRepo.Repository, uG Generator) *paymentUsecase {
	return &paymentUsecase{pRepository: pr, qRepository: qr, uuidGenerator: uG}
}
