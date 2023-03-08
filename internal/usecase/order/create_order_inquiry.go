package order

import (
	"context"
	"errors"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrCreateOrderInquiry_UserNotAuthorized = errors.New("CREATE_ORDER_INQUIRY.USER_NOT_AUTHORIZED")
)

func (u orderUsecase) CreateOrderInquiry(ctx context.Context, arg oModel.AddOrderInquiry, au uModel.AuthUser) (oModel.OrderInquiry, error) {
	if !au.IsSame(arg.UserID) {
		return oModel.OrderInquiry{}, ErrCreateOrderInquiry_UserNotAuthorized
	}

	// Calculate distance, Overview Polyline
	dr, err := u.mapCalculator.CalculateDirection(ctx, arg.Origin, arg.Destination)
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	/* Calculate price
	Rp2000 per km
	If price < 5000, make price = 5000
	*/

	arg.ID, err = u.uuidGenerator.GenerateUUID()
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	arg.Price = int64(dr.Distance) * 2
	arg.Duration = dr.Duration
	arg.Origin.Address = dr.Origin.Address
	arg.Destination.Address = dr.Destination.Address
	arg.Routes = dr.PolylineToStr()

	id, err := u.oRepository.CreateOrderInquiry(ctx, arg)
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	oi, err := u.oRepository.GetOrderInquiry(ctx, id)
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	return oi, nil
}
