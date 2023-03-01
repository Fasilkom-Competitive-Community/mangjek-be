package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	"github.com/gin-gonic/gin"
	"net/http"

	httpOrderCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (d HTTPOrderDelivery) addOrderInquiry(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	var oi httpOrderCommon.AddOrderInquiry
	if err := c.ShouldBindJSON(&oi); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	rdr, roi, err := d.orderUCase.CreateOrderInquiry(ctx, oModel.AddOrderInquiry{
		UserID: oi.UserID,
		Origin: oModel.Location{
			Latitude:  oi.Origin.Latitude,
			Longitude: oi.Origin.Longitude,
		},
		Destination: oModel.Location{
			Latitude:  oi.Destination.Latitude,
			Longitude: oi.Destination.Longitude,
		},
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, httpOrderCommon.OrderInquiry{
		ID:       roi.ID,
		UserID:   roi.UserID,
		Price:    roi.Price,
		Distance: roi.Distance,
		Duration: roi.Duration,
		Origin: httpOrderCommon.Location{
			Address:   roi.Origin.Address,
			Latitude:  roi.Origin.Latitude,
			Longitude: roi.Origin.Longitude,
		},
		Destination: httpOrderCommon.Location{
			Address:   roi.Destination.Address,
			Latitude:  roi.Destination.Latitude,
			Longitude: roi.Destination.Longitude,
		},
		CreatedAt: roi.CreatedAt,
		UpdatedAt: roi.UpdatedAt,
	})
}
