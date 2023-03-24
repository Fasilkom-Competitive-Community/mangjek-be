package order

import (
	"context"
	"errors"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrUpdateOrderStatus_UserNotAuthorized = errors.New("UPDATE_ORDER_STATUS.USER_NOT_AUTHORIZED")
)

func (u orderUsecase) UpdateOrderStatus(ctx context.Context, arg oModel.UpdateOrderStatus, au uModel.AuthUser) error {
	order, err := u.oRepository.GetOrder(ctx, arg.ID)
	if err != nil {
		return err
	}

	// check driverid
	if !au.IsSame(order.Driver.UserID) {
		return ErrUpdateOrderStatus_UserNotAuthorized
	}

	// check state (can't back to previous state)


	_, err = u.oRepository.UpdateOrderStatus(ctx, arg)
	return err
}
