package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	httpOrderCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPOrderDelivery) getOrderInquiry(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	inquiryId := c.Param("inquiryId")

	oi, err := d.orderUCase.GetOrderInquiry(ctx, inquiryId, au)
	if err != nil {
		c.Error(err)
		return
	}

	rList, err := oi.RoutesList()
	if err != nil {
		c.Error(err)
		return
	}

	resp := httpOrderCommon.OrderInquiry{
		ID:       oi.ID,
		UserID:   oi.UserID,
		Price:    oi.Price,
		Distance: oi.Distance,
		Duration: oi.Duration,
		Origin: httpOrderCommon.Location{
			Address:   oi.Origin.Address,
			Latitude:  oi.Origin.Latitude,
			Longitude: oi.Origin.Longitude,
		},
		Destination: httpOrderCommon.Location{
			Address:   oi.Destination.Address,
			Latitude:  oi.Destination.Latitude,
			Longitude: oi.Destination.Longitude,
		},
		CreatedAt: oi.CreatedAt,
		UpdatedAt: oi.UpdatedAt,
	}

	resp.Routes = make([]httpOrderCommon.Location, 0)
	for _, location := range rList {
		resp.Routes = append(resp.Routes, httpOrderCommon.Location{
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		})
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: resp})
}

func (d HTTPOrderDelivery) getOrder(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	uId := c.Param("ID")

	o, err := d.orderUCase.GetOrder(ctx, uId, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: o})
}
