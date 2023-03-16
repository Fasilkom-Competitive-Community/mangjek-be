package user

import (
	"context"
	"errors"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"strings"
)

var (
	ErrUpdateUser_UserNotAuthorized = errors.New("UPDATE_USER.USER_NOT_AUTHORIZED")
)

// UpdateUser implements Usecase
func (u userUsecase) UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", ErrUpdateUser_UserNotAuthorized
	}
	arg.Nim = strings.Split(arg.Email, "@")[0]
	return u.uRepository.UpdateUser(ctx, arg)
}
