package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	httpPaymentCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/payment"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPPaymentDelivery) paidQRISCallback(c *gin.Context) {
	ctx := c.Request.Context()
	paid, err := d.paymentUCase.UpdatePaymentStatusToPaid(ctx)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (d HTTPPaymentDelivery) addPayment(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	var p httpPaymentCommon.AddPayment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	payment, err := d.paymentUCase.CreatePayment(ctx, pModel.AddPayment{
		OrderID: "",
		Amount:  p.Amount,
		Method:  pModel.Method(p.Method),
	}, au)
	if err != nil {
		c.Error(err)
		return
	}
}
