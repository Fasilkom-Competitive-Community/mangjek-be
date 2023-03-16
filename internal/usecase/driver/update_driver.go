package user

import (
	"context"
	"errors"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrUpdateDriver_DriverNotAuthorizhed = errors.New("UPDATE_DRIVER.DRIVER_NOT_AUTHORIZED")
)

// UpdateDriver implements Usecase
func (u driverUsecase) UpdateDriver(ctx context.Context, arg dModel.UpdateDriver, au uModel.AuthUser) (int32, error) {
	if !au.IsSame(arg.UserID) {
		return 0, ErrUpdateDriver_DriverNotAuthorizhed
	}
	return u.dRepository.UpdateDriver(ctx, arg)
}
