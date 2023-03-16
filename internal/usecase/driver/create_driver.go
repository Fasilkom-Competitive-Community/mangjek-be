package user

import (
	"context"
	"errors"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrCreateDriver_UserNotAuthorized = errors.New("CREATE_DRIVER.DRIVER_NOT_AUTHORIZED")
	ErrCreateDriver_EmailNotVerified  = errors.New("CREATE_DRIVER.EMAIL_NOT_VERIFIED")
	ErrCreateDriver_DriverExist       = errors.New("CREATE_DRIVER.DRIVER_EXISTS")
)

// CreateDriver implements Usecase
func (u driverUsecase) CreateDriver(ctx context.Context, arg dModel.AddDriver, au uModel.AuthUser) (int32, error) {
	if !au.IsSame(arg.UserID) {
		return 0, ErrCreateDriver_UserNotAuthorized
	}

	auf, err := u.auRepository.GetAuthUserFull(ctx, au.ID)
	if err != nil {
		return 0, err
	}

	if !auf.IsEmailVerified {
		return 0, ErrCreateDriver_EmailNotVerified
	}

	// check if driver already exists
	available, err := u.dRepository.VerifyAvailableDriver(ctx, arg.UserID)
	if err != nil {
		return 0, err
	}

	if available {
		return 0, ErrCreateDriver_DriverExist
	}

	return u.dRepository.CreateDriver(ctx, arg)
}
