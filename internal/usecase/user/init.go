package user

import (
	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user"
)

type userUsecase struct {
	uRepository  uRepo.Repository
	auRepository auRepo.Repository
}

func NewUserUsecase(ur uRepo.Repository, aur auRepo.Repository) Usecase {
	return userUsecase{uRepository: ur, auRepository: aur}
}
