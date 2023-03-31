package http

import (
	"fmt"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	httpOrderCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

	oId := c.Param("orderId")
	fmt.Println(oId)

	o, err := d.orderUCase.GetOrder(ctx, oId, au)
	if err != nil {
		c.Error(err)
		return
	}

	resp := httpOrderCommon.GetOrder{
		ID:    oId,
		UName: o.User.Name,
		Driver: httpOrderCommon.Driver{
			Name:         o.DName,
			PoliceNumber: o.Driver.PoliceNumber,
			VehicleType:  o.Driver.VehicleType,
			VehicleModel: o.Driver.VehicleModel,
		},
		OrderInquiry: httpOrderCommon.Inquiry{
			Price:    o.OrderInquiry.Price,
			Distance: o.OrderInquiry.Distance,
			Duration: o.OrderInquiry.Duration,
			OAddress: o.OrderInquiry.Origin.Address,
			DAddress: o.OrderInquiry.Destination.Address,
		},
		Payment: httpOrderCommon.Payment{
			Amount:   o.Payment.Amount,
			Status:   string(o.Payment.Status),
			Method:   string(o.Payment.Method),
			QrString: o.Payment.QrString,
		},
		Status:    string(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: resp})
}

func (d HTTPOrderDelivery) getOrderRoutes(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	oId := c.Param("orderId")
	fmt.Println(oId)

	o, err := d.orderUCase.GetOrder(ctx, oId, au)
	if err != nil {
		c.Error(err)
		return
	}

	routes := o.OrderInquiry.Routes

	routesArr := strings.Split(routes, ";")

	var locs []httpOrderCommon.Location
	for i := 0; i < len(routesArr); i++ {
		temp := strings.Split(routesArr[i], ",")

		lat, err := strconv.ParseFloat(temp[0], 64)
		if err != nil {
			panic(err)
		}
		long, err := strconv.ParseFloat(temp[1], 64)
		if err != nil {
			panic(err)
		}

		locs = append(locs, httpOrderCommon.Location{
			Latitude:  lat,
			Longitude: long,
		})
	}

	resp := httpOrderCommon.GetOrderRoutes{
		ID:     oId,
		Routes: locs,
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: resp})
}
