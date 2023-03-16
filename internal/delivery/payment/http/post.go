package http

import (
	"github.com/gin-gonic/gin"
)

func (d HTTPPaymentDelivery) paidQRISCallback(c *gin.Context) {
	// ctx := c.Request.Context()
	// paid, err := d.paymentUCase.UpdatePaymentStatusToPaid(ctx)
	// if err != nil {
	// 	c.Error(err)
	// 	return
	// }
	// c.JSON(http.StatusOK, "ok")
	return
}

func (d HTTPPaymentDelivery) addPayment(c *gin.Context) {
	// ctx := c.Request.Context()
	// au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	// var p httpPaymentCommon.AddPayment
	// if err := c.ShouldBindJSON(&p); err != nil {
	// 	c.Error(err).SetType(gin.ErrorTypeBind)
	// 	return
	// }

	// payment, err := d.paymentUCase.CreatePayment(ctx, pModel.AddPayment{
	// 	OrderID: "",
	// 	Amount:  p.Amount,
	// 	Method:  pModel.Method(p.Method),
	// }, au)
	// if err != nil {
	// 	c.Error(err)
	// 	return
	// }
	return
}
