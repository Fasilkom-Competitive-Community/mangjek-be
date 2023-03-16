package order

import (
	"context"
	"errors"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrCreateOrder_UserNotAuthorized     = errors.New("CREATE_ORDER.USER_NOT_AUTHORIZED")
	ErrCreateOrder_DriverIsSameAsUser    = errors.New("CREATE_ORDER.DRIVER_IS_SAME_AS_USER")
	ErrCreateOrder_PaymentMethodNotFound = errors.New("CREATE_ORDER.PAYMENT_METHOD_NOT_FOUND")
)

func (u orderUsecase) CreateOrder(ctx context.Context, arg oModel.AddOrder, au uModel.AuthUser) (oModel.Order, error) {
	// Check userid
	if !au.IsSame(arg.UserID) {
		return oModel.Order{}, ErrCreateOrder_UserNotAuthorized
	}

	// Check if driver is same as user
	us, err := u.uRepository.GetUser(ctx, arg.UserID)
	if err != nil {
		return oModel.Order{}, err
	}
	d, err := u.dRepository.GetDriver(ctx, arg.DriverID)
	if err != nil {
		return oModel.Order{}, err
	}
	if d.UserID == us.ID {
		return oModel.Order{}, ErrCreateOrder_DriverIsSameAsUser
	}

	// Check order inquiry userid
	oi, err := u.oRepository.GetOrderInquiry(ctx, arg.OrderInquiryID)
	if err != nil {
		return oModel.Order{}, err
	}
	if !au.IsSame(oi.UserID) {
		return oModel.Order{}, ErrCreateOrder_UserNotAuthorized
	}

	// Generate order
	pid, err := u.uuidGenerator.GenerateUUID()
	if err != nil {
		return oModel.Order{}, err
	}
	ap := pModel.AddPayment{
		ID:     pid,
		Amount: arg.Payment.Amount,
		Method: arg.Payment.Method,
	}
	switch ap.Method {
	case "CASH":
		ap.Status = pModel.PaidStatus
	case "QRIS":
		ap.Status = pModel.UnpaidStatus
	default:
		return oModel.Order{}, ErrCreateOrder_PaymentMethodNotFound
	}

	oid, err := u.uuidGenerator.GenerateUUID()
	if err != nil {
		return oModel.Order{}, err
	}
	ao := oModel.AddOrder{
		ID:             oid,
		UserID:         us.ID,
		DriverID:       d.ID,
		OrderInquiryID: oi.ID,
		Payment:        ap,
		Status:         oModel.OnProgressStatus,
	}
	_, err = u.oRepository.CreateOrder(ctx, ap, ao)
	if err != nil {
		return oModel.Order{}, err
	}

	o := oModel.Order{
		ID: oid,
		User: uModel.User{
			ID:          us.ID,
			Name:        us.Name,
			Email:       us.Email,
			PhoneNumber: us.PhoneNumber,
			Nim:         us.Nim,
		},
		Driver: dModel.Driver{
			ID:           d.ID,
			UserID:       d.UserID,
			PoliceNumber: d.PoliceNumber,
			VehicleModel: d.VehicleModel,
			VehicleType:  d.VehicleType,
		},
		OrderInquiry: oModel.OrderInquiry{
			ID:           oi.ID,
			UserID:       oi.UserID,
			OrderRouteID: oi.OrderRouteID,
			Price:        oi.Price,
			Distance:     oi.Distance,
			Duration:     oi.Duration,
			Origin:       oi.Origin,
			Destination:  oi.Destination,
			Routes:       oi.Routes,
		},
		Payment: pModel.Payment{
			ID:     pid,
			Amount: ap.Amount,
			Status: ap.Status,
			Method: ap.Method,
		},
		Status: ao.Status,
	}

	return o, nil
}
