package pg

import (
	"context"
	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	oRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/order"
	"github.com/jackc/pgx/v4"
)

type pgOrderInquiryRepository struct {
	querier sqlc.Querier
}

// CreateOrderInquiry implements order.OrderInquiry
func (r pgOrderInquiryRepository) CreateOrderInquiry(ctx context.Context, arg oModel.AddOrderInquiry) (string, error) {
	id, err := r.querier.CreateOrderInquiry(ctx, sqlc.CreateOrderInquiryParams{
		ID:                 arg.ID,
		UserID:             arg.UserID,
		Price:              arg.Price,
		Distance:           arg.Distance,
		Duration:           arg.Duration,
		OriginLat:          arg.Origin.Latitude,
		OriginLong:         arg.Origin.Longitude,
		OriginAddress:      arg.Origin.Address,
		DestinationLat:     arg.Destination.Latitude,
		DestinationLong:    arg.Destination.Longitude,
		DestinationAddress: arg.Destination.Address,
		Routes:             arg.Routes,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("OrderInquiry not found")
	}
	return id, err
}

// DeleteOrderInquiry implements order.OrderInquiry
func (r pgOrderInquiryRepository) DeleteOrderInquiry(ctx context.Context, id string) error {
	err := r.querier.DeleteOrderInquiry(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("OrderInquiry not found")
	}
	return err
}

// GetOrderInquiry implements order.OrderInquiry
func (r pgOrderInquiryRepository) GetOrderInquiry(ctx context.Context, id string) (oModel.OrderInquiry, error) {
	o, err := r.querier.GetOrderInquiry(ctx, id)
	if err == pgx.ErrNoRows {
		return oModel.OrderInquiry{}, errorCommon.NewNotFoundError("OrderInquiry not found")
	}

	return oModel.OrderInquiry{
		ID:       o.ID,
		UserID:   o.UserID,
		Price:    o.Price,
		Distance: o.Distance,
		Duration: o.Duration,
		Origin: oModel.Location{
			Address:   o.OriginAddress,
			Latitude:  o.OriginLat,
			Longitude: o.OriginLong,
		},
		Destination: oModel.Location{
			Address:   o.DestinationAddress,
			Latitude:  o.DestinationLat,
			Longitude: o.DestinationLong,
		},
		Routes:    o.Routes,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}, nil
}

func NewPGOrderInquiryRepository(querier sqlc.Querier) oRepo.Repository {
	return pgOrderInquiryRepository{querier: querier}
}
