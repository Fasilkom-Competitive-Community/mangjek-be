package pg

import (
	"context"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"

	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	uRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/user"
	"github.com/jackc/pgx/v4"
)

type pgUserRepository struct {
	querier sqlc.Querier
}

// CreateUser implements user.User
func (r pgUserRepository) CreateUser(ctx context.Context, arg uModel.AddUser) (string, error) {
	id, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams(arg))
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
	return uModel.User(u), err
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
		ums = append(ums, uModel.User(u))
	}
	return ums, err
}

// UpdateUser implements user.User
func (r pgUserRepository) UpdateUser(ctx context.Context, arg uModel.UpdateUser) (string, error) {
	id, err := r.querier.UpdateUser(ctx, sqlc.UpdateUserParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

func NewPGUserRepository(querier sqlc.Querier) uRepo.Repository {
	return pgUserRepository{querier: querier}
}

// GetUserHistory implements BLM
func (r pgUserRepository) GetUserHistory(ctx context.Context, id string) ([]oModel.Order, error) {
	o, err := r.querier.GetOrderHistory(ctx, id)
	if err == pgx.ErrNoRows {
		var temp []oModel.Order
		return temp, errorCommon.NewNotFoundError("User not found")
	}

	var oh []oModel.Order

	for i := 0; i < len(o); i++ {
		d, err := r.querier.GetUser(ctx, o[i].UserID_2)
		if err == pgx.ErrNoRows {
			var temp []oModel.Order
			return temp, errorCommon.NewNotFoundError("Driver not found")
		}

		oh = append(oh, oModel.Order{
			ID:    o[i].ID,
			DName: d.Name,
			User: uModel.User{
				ID:   o[i].UserID,
				Name: o[i].Name,
			},
			Driver: dModel.Driver{
				ID:           o[i].DriverID,
				PoliceNumber: o[i].PoliceNumber,
				VehicleModel: o[i].VehicleModel,
				VehicleType:  o[i].VehicleType,
			},
			OrderInquiry: oModel.OrderInquiry{
				ID:       o[i].OrderInquiryID,
				Price:    o[i].Price,
				Distance: o[i].Distance,
				Duration: o[i].Duration,
				Origin: oModel.Location{
					Address: o[i].OriginAddress,
				},
				Destination: oModel.Location{
					Address: o[i].DestinationAddress,
				},
				Routes: o[i].Routes,
			},
			Payment: pModel.Payment{
				ID:       o[i].PaymentID,
				Amount:   o[i].Amount,
				Status:   pModel.Status(o[i].Status_2),
				Method:   pModel.Method(o[i].Method),
				QrString: o[i].QrStr,
			},
			Status:    oModel.Status(o[i].Status),
			CreatedAt: o[i].CreatedAt,
			UpdatedAt: o[i].UpdatedAt,
		})
	}

	return []oModel.Order(oh), err

	//{
	//	"id" : 1,
	//	"address" : "Gang Buntu",
	//	"status" : "Sedang diperjalanan",
	//	"update_at" : "2023-02-10T13:45:00.000Z"
	//
	//}
}
