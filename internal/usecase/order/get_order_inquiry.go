package order

import (
	"context"
	"errors"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

var (
	ErrGetOrderInquiry_UserNotAuthorized = errors.New("GET_ORDER_INQUIRY.USER_NOT_AUTHORIZED")
)

func (u orderUsecase) GetOrderInquiry(ctx context.Context, id string, au uModel.AuthUser) (oModel.OrderInquiry, error) {
	oi, err := u.oRepository.GetOrderInquiry(ctx, id)
	if err != nil {
		return oModel.OrderInquiry{}, err
	}

	if !au.IsSame(oi.UserID) {
		return oModel.OrderInquiry{}, ErrGetOrderInquiry_UserNotAuthorized
	}

	return oi, nil
}
