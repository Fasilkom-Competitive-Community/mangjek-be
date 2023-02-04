package user

import (
	"context"
	"errors"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	auRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/auth"
	dRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"
)

type driverUsecase struct {
	dRepository  dRepo.Repository
	auRepository auRepo.Repository
}

// CreateDriver implements Usecase
func (u driverUsecase) CreateDriver(ctx context.Context, arg dModel.AddDriver, au uModel.AuthUser) (int32, error) {
	if !au.IsSame(arg.UserID) {
		return 0, errors.New("CREATE_DRIVER.DRIVER_NOT_AUTHORIZED")
	}

	auf, err := u.auRepository.GetAuthUserFull(ctx, au.ID)
	if err != nil {
		return 0, err
	}

	if !auf.IsEmailVerified {
		return 0, errors.New("CREATE_DRIVER.EMAIL_NOT_VERIFIED")
	}

	// check if driver already exists
	available, err := u.dRepository.VerifyAvailableDriver(ctx, arg.UserID)
	if err != nil {
		return 0, err
	}

	if available {
		return 0, errors.New("CREATE_DRIVER.DRIVER_EXISTS")
	}

	return u.dRepository.CreateDriver(ctx, arg)
}

// DeleteDriver implements Usecase
//
// # Not implemented for now
//
// -
func (u driverUsecase) DeleteDriver(ctx context.Context, id int32, uid string, au uModel.AuthUser) error {
	return nil
	// return u.dRepository.DeleteDriver(ctx, id)
}

// GetDriver implements Usecase
func (u driverUsecase) GetDriver(ctx context.Context, id int32, au uModel.AuthUser) (dModel.Driver, error) {
	driver, err := u.dRepository.GetDriver(ctx, id)
	if err == nil {
		if au.IsSame(driver.UserID) || au.IsAdmin() {
			return driver, nil
		}
		return dModel.Driver{}, errors.New("GET_DRIVER.DRIVER_NOT_AUTHORIZED")
	}
	return driver, err
}

// ListDrivers implements Usecase
func (u driverUsecase) ListDrivers(ctx context.Context, au uModel.AuthUser) ([]dModel.Driver, error) {
	if !au.IsAdmin() {
		return nil, errors.New("LIST_DRIVER.DRIVER_NOT_AUTHORIZED")
	}
	return u.dRepository.ListDrivers(ctx)
}

// UpdateDriver implements Usecase
func (u driverUsecase) UpdateDriver(ctx context.Context, arg dModel.UpdateDriver, au uModel.AuthUser) (int32, error) {
	panic("not implemented")
}

func NewDriverUsecase(dr dRepo.Repository, aur auRepo.Repository) Usecase {
	return driverUsecase{dRepository: dr, auRepository: aur}
}
