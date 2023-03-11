package order

import (
	"context"
	"errors"
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

	return oModel.Order{}, nil
}
