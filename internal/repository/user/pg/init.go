package pg

import (
	"context"

	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	sRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user"
	"github.com/jackc/pgx/v4"
)

type pgUserRepository struct {
	querier sqlc.Querier
}

// CreateUser implements user.User
func (r pgUserRepository) CreateUser(ctx context.Context, arg uModel.AddUser) (string, error) {
	id, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams{
		ID: arg.ID,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

// DeleteUser implements user.User
func (r pgUserRepository) DeleteUser(ctx context.Context, id string) error {
	err := r.querier.DeleteUser(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("User not found")
	}
	return err
}

// GetUser implements user.User
func (r pgUserRepository) GetUser(ctx context.Context, id string) (uModel.User, error) {
	u, err := r.querier.GetUser(ctx, id)
	if err == pgx.ErrNoRows {
		return uModel.User{}, errorCommon.NewNotFoundError("User not found")
	}
	return uModel.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, err
}

// VerifyAvailableUser implements user.User
func (r pgUserRepository) VerifyAvailableUser(ctx context.Context, id string) (bool, error) {
	u, err := r.querier.GetUser(ctx, id)
	// user not available
	if err == pgx.ErrNoRows || (err == nil && u.ID != id) {
		return false, nil
	}
	// error
	if err != nil {
		return false, err
	}
	// user available
	return true, nil
}

// ListUsers implements user.User
func (r pgUserRepository) ListUsers(ctx context.Context) ([]uModel.User, error) {
	us, err := r.querier.ListUsers(ctx)
	ums := make([]uModel.User, 0)
	for _, u := range us {
		ums = append(ums, uModel.User{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}
	return ums, err
}

// UpdateUser implements user.User
func (r pgUserRepository) UpdateUser(ctx context.Context, arg uModel.UpdateUser) (string, error) {
	id, err := r.querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:    arg.ID,
		Name:  arg.Name,
		Email: arg.Email,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

func NewPGUserRepository(querier sqlc.Querier) sRepo.Repository {
	return pgUserRepository{querier: querier}
}
