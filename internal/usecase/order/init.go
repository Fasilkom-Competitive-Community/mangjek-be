package order

import (
	dRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"
	oRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/order"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user"
)

type orderUsecase struct {
	oRepository oRepo.Repository
	uRepository uRepo.Repository
	dRepository dRepo.Repository
	
	mapCalculator Calculator
	uuidGenerator Generator
}

func NewOrderUsecase(or oRepo.Repository, ur uRepo.Repository, dr dRepo.Repository, mC Calculator, uG Generator) *orderUsecase {
	return &orderUsecase{oRepository: or, uRepository: ur, dRepository: dr, mapCalculator: mC, uuidGenerator: uG}
}
