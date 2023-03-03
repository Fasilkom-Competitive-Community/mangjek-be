package http

import (
	"firebase.google.com/go/v4/auth"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	pUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/payment"
	"github.com/gin-gonic/gin"
)

type HTTPPaymentDelivery struct {
	paymentUCase pUCase.Usecase
}

func NewHTTPPaymentDelivery(g *gin.RouterGroup, paymentUCase pUCase.Usecase, fAuth *auth.Client) HTTPPaymentDelivery {
	h := HTTPPaymentDelivery{paymentUCase: paymentUCase}

	g.POST("/payments", http.Auth(fAuth), h.addPayment)
	g.POST("/payments/qris/callback", h.paidQRISCallback)

	return h
}
