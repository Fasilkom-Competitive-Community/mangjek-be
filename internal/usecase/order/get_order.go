package order

import (
	"context"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (u orderUsecase) GetOrder(ctx context.Context, id string, au uModel.AuthUser) (oModel.Order, error) {
	o, err := u.oRepository.GetOrder(ctx, id)
	if err != nil {
		return oModel.Order{}, err
	}

	if !au.IsSame(o.User.ID) {
		return oModel.Order{}, ErrGetOrderInquiry_UserNotAuthorized
	}

	return o, err
}
