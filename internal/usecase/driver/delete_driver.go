package user

import (
	"context"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

// DeleteDriver implements Usecase
//
// # Not implemented for now
//
// -
func (u driverUsecase) DeleteDriver(ctx context.Context, id int32, uid string, au uModel.AuthUser) error {
	return nil
	// return u.dRepository.DeleteDriver(ctx, id)
}
