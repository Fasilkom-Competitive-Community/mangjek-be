package user

import (
	"context"
	"errors"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"strings"
)

var (
	ErrCreateUser_UserNotAuthorized = errors.New("CREATE_USER.USER_NOT_AUTHORIZED")
	ErrCreateUser_UserExist         = errors.New("CREATE_USER.USER_EXISTS")
)

func (u userUsecase) CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", ErrCreateUser_UserNotAuthorized
	}

	auf, err := u.auRepository.GetAuthUserFull(ctx, au.ID)
	if err != nil {
		return "", err
	}

	if !auf.IsEmailVerified {
		return "", ErrCreateUser_UserNotAuthorized
	}

	// check if user already exists
	available, err := u.uRepository.VerifyAvailableUser(ctx, au.ID)
	if err != nil {
		return "", err
	}

	if available {
		return "", ErrCreateUser_UserExist
	}

	arg.Nim = strings.Split(arg.Email, "@student.unsri.ac.id")[0]
	return u.uRepository.CreateUser(ctx, arg)
}
