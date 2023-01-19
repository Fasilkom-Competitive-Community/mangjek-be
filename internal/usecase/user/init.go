package user

import (
	"context"
	"errors"
	"strings"

	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user"
)

type userUsecase struct {
	uRepository  uRepo.Repository
	auRepository auRepo.Repository
}

// CreateUser implements Usecase
func (u userUsecase) CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", errors.New("CREATE_USER.USER_NOT_AUTHORIZED")
	}

	auf, err := u.auRepository.GetAuthUserFull(ctx, au.ID)
	if err != nil {
		return "", err
	}

	if !auf.IsEmailVerified {
		return "", errors.New("CREATE_USER.EMAIL_NOT_VERIFIED")
	}

	// check if user already exists
	available, err := u.uRepository.VerifyAvailableUser(ctx, au.ID)
	if err != nil {
		return "", err
	}

	if available {
		return "", errors.New("CREATE_USER.USER_EXISTS")
	}

	arg.NIM = strings.Split(arg.Email, "@")[0]
	return u.uRepository.CreateUser(ctx, arg)
}

// DeleteUser implements Usecase
//
// # Not implemented for now
//
// -
func (u userUsecase) DeleteUser(ctx context.Context, id string, au uModel.AuthUser) error {
	return nil
	// return u.uRepository.DeleteUser(ctx, id)
}

// GetUser implements Usecase
func (u userUsecase) GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error) {
	if au.IsSame(id) || au.IsAdmin() {
		return u.uRepository.GetUser(ctx, id)
	}
	return uModel.User{}, errors.New("GET_USER.USER_NOT_AUTHORIZED")
}

// ListUsers implements Usecase
func (u userUsecase) ListUsers(ctx context.Context, au uModel.AuthUser) ([]uModel.User, error) {
	if !au.IsAdmin() {
		return nil, errors.New("LIST_USER.USER_NOT_AUTHORIZED")
	}
	return u.uRepository.ListUsers(ctx)
}

// UpdateUser implements Usecase
func (u userUsecase) UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", errors.New("UPDATE_USER.USER_NOT_AUTHORIZED")
	}
	arg.NIM = strings.Split(arg.Email, "@")[0]
	return u.uRepository.UpdateUser(ctx, arg)
}

func NewUserUsecase(ur uRepo.Repository, aur auRepo.Repository) Usecase {
	return userUsecase{uRepository: ur, auRepository: aur}
}
