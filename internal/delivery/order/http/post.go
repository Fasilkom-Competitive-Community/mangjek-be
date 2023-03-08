package http

import (
	"fmt"
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

	fmt.Println("s,ssmm")

	var oi httpOrderCommon.AddOrderInquiry
	if err := c.ShouldBindJSON(&oi); err != nil {
		fmt.Println(err)
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	fmt.Println("ini io", oi)

	roi, err := d.orderUCase.CreateOrderInquiry(ctx, oModel.AddOrderInquiry{
		UserID: oi.UserID,
		Origin: oModel.Location{
			Latitude:  oi.OriginLatitude,
			Longitude: oi.OriginLongitude,
		},
		Destination: oModel.Location{
			Latitude:  oi.DestinationLatitude,
			Longitude: oi.DestinationLongitude,
		},
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	fmt.Println("ini roi", oi)

	rList, err := roi.RoutesList()
	if err != nil {
		c.Error(err)
		return
	}

	resp := httpOrderCommon.OrderInquiry{
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
	}

	resp.Routes = make([]httpOrderCommon.Location, 0)
	for _, location := range rList {
		resp.Routes = append(resp.Routes, httpOrderCommon.Location{
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		})
	}

	c.JSON(http.StatusCreated, resp)
}
