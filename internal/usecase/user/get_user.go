package user

import (
	"context"
	"errors"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrGetUser_UserNotAuthorized   = errors.New("GET_USER.USER_NOT_AUTHORIZED")
	ErrListUseer_UserNotAuthorized = errors.New("LIST_USER.USER_NOT_AUTHORIZED")
)

func (u userUsecase) GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error) {
	if au.IsSame(id) || au.IsAdmin() {
		return u.uRepository.GetUser(ctx, id)
	}
	return uModel.User{}, ErrGetUser_UserNotAuthorized
}

// ListUsers implements Usecase
func (u userUsecase) ListUsers(ctx context.Context, au uModel.AuthUser) ([]uModel.User, error) {
	if !au.IsAdmin() {
		return nil, ErrListUseer_UserNotAuthorized
	}
	return u.uRepository.ListUsers(ctx)
}
