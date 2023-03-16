package user

import (
	"context"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

// not implemented for now
func (u userUsecase) DeleteUser(ctx context.Context, id string, au uModel.AuthUser) error {
	return nil
	// return u.uRepository.DeleteUser(ctx, id)
}
