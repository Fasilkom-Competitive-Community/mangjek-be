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

	arg.ID, err = u.uuidGenerator.GenerateUUID()
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	/* Calculate price
	<= 3km → Rp 5.000
	> 3km, per kilo dihitung Rp2.000
	ex: jarak 3.2 km = Rp5.000 + (0.2 * Rp2.000) = Rp5.400 dibulatkan ke atas (pecahan Rp500) menjadi Rp5.500
	*/

	if dr.Distance <= 3_000 {
		arg.Price = 5_000
	} else {
		arg.Price = int64(dr.Distance) * 2
		remainder := arg.Price % 500
		if remainder != 0 {
			arg.Price = arg.Price - remainder + 500
		}
	}

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
