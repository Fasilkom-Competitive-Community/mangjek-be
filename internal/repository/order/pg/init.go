package pg

import (
	"context"
	errorCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/error"
	store "github.com/Fasilkom-Competitive-Community/mangjek-be/common/pg"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/sqlc"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	oRepo "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/repository/order"
	"github.com/jackc/pgx/v4"
)

type pgOrderInquiryRepository struct {
	querier *store.Store
}

// CreateOrderInquiry implements order.Order
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

// DeleteOrderInquiry implements order.Order
func (r pgOrderInquiryRepository) DeleteOrderInquiry(ctx context.Context, id string) error {
	err := r.querier.DeleteOrderInquiry(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("OrderInquiry not found")
	}
	return err
}

// GetOrderInquiry implements order.Order
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

// CreateOrder implements order.Order
func (r pgOrderInquiryRepository) CreateOrder(ctx context.Context, payment pModel.AddPayment, order oModel.AddOrder) (string, error) {
	var oid string
	err := r.querier.ExecTx(ctx, func(q sqlc.Querier) error {
		pid, err := q.CreatePayment(ctx, sqlc.CreatePaymentParams{
			ID:     payment.ID,
			Amount: payment.Amount,
			Status: string(payment.Status),
			Method: string(payment.Method),
			QrStr:  payment.QrString,
		})
		if err != nil {
			return err
		}

		order.Payment.ID = pid
		oid, err = q.CreateOrder(ctx, sqlc.CreateOrderParams{
			ID:             order.ID,
			UserID:         order.UserID,
			DriverID:       order.DriverID,
			OrderInquiryID: order.OrderInquiryID,
			PaymentID:      order.Payment.ID,
			Status:         string(order.Status),
		})
		return err
	})
	if err != nil {
		return "", err
	}

	return oid, nil
}

func NewPGOrderInquiryRepository(querier *store.Store) oRepo.Repository {
	return pgOrderInquiryRepository{querier: querier}
}
