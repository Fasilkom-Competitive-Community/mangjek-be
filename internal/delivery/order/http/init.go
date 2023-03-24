package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	oUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/order"
	"github.com/gin-gonic/gin"
)

type HTTPOrderDelivery struct {
	orderUCase oUCase.Usecase
}

func NewHTTPOrderDelivery(g *gin.RouterGroup, orderUCase oUCase.Usecase, fAuth *auth.Client) HTTPOrderDelivery {
	h := HTTPOrderDelivery{orderUCase: orderUCase}

	g.POST("/orders", httpCommon.Auth(fAuth), h.addOrder)
	//g.PATCH("/orders/:id", httpCommon.Auth(fAuth), h.updateOrderStatus)
	g.GET("/orders/inquiry/:inquiryId", httpCommon.Auth(fAuth), h.getOrderInquiry)
	g.POST("/orders/inquiry", httpCommon.Auth(fAuth), h.addOrderInquiry)

	return h
}
