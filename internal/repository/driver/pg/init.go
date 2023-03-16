package pg

import (
	"context"

	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	dRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"
	"github.com/jackc/pgx/v4"
)

type pgDriverRepository struct {
	querier sqlc.Querier
}

// CreateDriver implements driver.Driver
func (r pgDriverRepository) CreateDriver(ctx context.Context, arg dModel.AddDriver) (int32, error) {
	id, err := r.querier.CreateDriver(ctx, sqlc.CreateDriverParams(arg))
	if err == pgx.ErrNoRows {
		return 0, errorCommon.NewNotFoundError("Driver not found")
	}
	return id, err
}

// DeleteDriver implements driver.Driver
func (r pgDriverRepository) DeleteDriver(ctx context.Context, id int32) error {
	err := r.querier.DeleteDriver(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("Driver not found")
	}
	return err
}

// GetDriver implements driver.Driver
func (r pgDriverRepository) GetDriver(ctx context.Context, id int32) (dModel.Driver, error) {
	u, err := r.querier.GetDriver(ctx, id)
	if err == pgx.ErrNoRows {
		return dModel.Driver{}, errorCommon.NewNotFoundError("Driver not found")
	}
	return dModel.Driver(u), err
}

// GetDriver implements driver.Driver
func (r pgDriverRepository) GetDriverByUserID(ctx context.Context, uid string) (dModel.Driver, error) {
	u, err := r.querier.GetDriverByUserID(ctx, uid)
	if err == pgx.ErrNoRows {
		return dModel.Driver{}, errorCommon.NewNotFoundError("Driver not found")
	}
	return dModel.Driver(u), err
}

// VerifyAvailableDriver implements driver.Driver
func (r pgDriverRepository) VerifyAvailableDriver(ctx context.Context, uid string) (bool, error) {
	u, err := r.querier.GetDriverByUserID(ctx, uid)
	// driver not available
	if err == pgx.ErrNoRows || (err == nil && u.UserID != uid) {
		return false, nil
	}
	// error
	if err != nil {
		return false, err
	}
	// driver available
	return true, nil
}

// ListDrivers implements driver.Driver
func (r pgDriverRepository) ListDrivers(ctx context.Context) ([]dModel.Driver, error) {
	us, err := r.querier.ListDrivers(ctx)
	ums := make([]dModel.Driver, 0)
	for _, u := range us {
		ums = append(ums, dModel.Driver(u))
	}
	return ums, err
}

// UpdateDriver implements driver.Driver
func (r pgDriverRepository) UpdateDriver(ctx context.Context, arg dModel.UpdateDriver) (int32, error) {
	id, err := r.querier.UpdateDriver(ctx, sqlc.UpdateDriverParams(arg))
	if err == pgx.ErrNoRows {
		return 0, errorCommon.NewNotFoundError("Driver not found")
	}
	return id, err
}

func NewPGDriverRepository(querier sqlc.Querier) dRepo.Repository {
	return pgDriverRepository{querier: querier}
}
