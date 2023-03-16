package user

import (
	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth"
	dRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"
)

type driverUsecase struct {
	dRepository  dRepo.Repository
	auRepository auRepo.Repository
}

func NewDriverUsecase(dr dRepo.Repository, aur auRepo.Repository) Usecase {
	return driverUsecase{dRepository: dr, auRepository: aur}
}
