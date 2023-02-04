package pg

import (
	"context"
	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	sRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"

	//sRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/driver"
	"github.com/jackc/pgx/v4"
)

type pgDriverRepository struct {
	querier sqlc.Querier
}

func (r pgDriverRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

// CreateDriver implements Driver.Driver
func (r pgDriverRepository) CreateDriver(ctx context.Context, arg uModel.AddDriver) (string, error) {
	id, err := r.querier.CreateDriver(ctx, sqlc.CreateDriverParams{
		ID: string(arg.ID),
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Driver not found")
	}
	return id, err
}

// DeleteDriver implements Driver.Driver
func (r pgDriverRepository) DeleteDriver(ctx context.Context, id string) error {
	err := r.querier.DeleteDriver(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("Driver not found")
	}
	return err
}

// GetDriver implements Driver.Driver
func (r pgDriverRepository) GetDriver(ctx context.Context, id string) (uModel.Driver, error) {
	u, err := r.querier.GetDriver(ctx, id)
	if err == pgx.ErrNoRows {
		return uModel.Driver{}, errorCommon.NewNotFoundError("Driver not found")
	}
	return uModel.Driver{
		ID:        u.ID,
		Nama:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, err
}

// VerifyAvailableDriver implements Driver.Driver
func (r pgDriverRepository) VerifyAvailableDriver(ctx context.Context, id string) (bool, error) {
	u, err := r.querier.GetDriver(ctx, id)
	// Driver not available
	if err == pgx.ErrNoRows || (err == nil && u.ID != id) {
		return false, nil
	}
	// error
	if err != nil {
		return false, err
	}
	// Driver available
	return true, nil
}

// ListDrivers implements Driver.Driver
func (r pgDriverRepository) ListDrivers(ctx context.Context) ([]uModel.Driver, error) {
	us, err := r.querier.ListDrivers(ctx)
	ums := make([]uModel.Driver, 0)
	for _, u := range us {
		ums = append(ums, uModel.Driver{
			ID:        u.ID,
			Nama:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return ums, err
}

// UpdateDriver implements Driver.Driver
func (r pgDriverRepository) UpdateDriver(ctx context.Context, arg uModel.UpdateDriver) (string, error) {
	id, err := r.querier.UpdateDriver(ctx, sqlc.UpdateDriverParams{
		ID:    string(arg.ID),
		Name:  arg.Nama,
		Email: arg.Email,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Driver not found")
	}
	return id, err
}

func NewPGDriverRepository(querier sqlc.Querier) sRepo.Repository {
	return pgDriverRepository{querier: querier}
}
