package order

import (
	oRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/order"
)

type orderUsecase struct {
	oRepository   oRepo.Repository
	mapCalculator Calculator
	uuidGenerator Generator
}

func NewOrderUsecase(or oRepo.Repository, mC Calculator, uG Generator) *orderUsecase {
	return &orderUsecase{oRepository: or, mapCalculator: mC, uuidGenerator: uG}
}
