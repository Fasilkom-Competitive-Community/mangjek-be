package user

import (
	"context"
	"errors"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrGetDriver_DriverNotAuthorized  = errors.New("GET_DRIVER.DRIVER_NOT_AUTHORIZED")
	ErrListDriver_DriverNotAuthorized = errors.New("LIST_DRIVER.DRIVER_NOT_AUTHORIZED")
)

// GetDriver implements Usecase
func (u driverUsecase) GetDriver(ctx context.Context, id int32, au uModel.AuthUser) (dModel.Driver, error) {
	driver, err := u.dRepository.GetDriver(ctx, id)
	if err == nil {
		if au.IsSame(driver.UserID) || au.IsAdmin() {
			return driver, nil
		}
		return dModel.Driver{}, ErrGetDriver_DriverNotAuthorized
	}
	return driver, err
}

// ListDrivers implements Usecase
func (u driverUsecase) ListDrivers(ctx context.Context, au uModel.AuthUser) ([]dModel.Driver, error) {
	if !au.IsAdmin() {
		return nil, ErrListDriver_DriverNotAuthorized
	}
	return u.dRepository.ListDrivers(ctx)
}
