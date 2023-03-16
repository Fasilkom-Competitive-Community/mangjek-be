package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	dUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/driver"
	"github.com/gin-gonic/gin"
)

type HTTPDriverDelivery struct {
	driverUCase dUCase.Usecase
}

func NewHTTPDriverDelivery(g *gin.RouterGroup, driverUCase dUCase.Usecase, fAuth *auth.Client) HTTPDriverDelivery {
	h := HTTPDriverDelivery{driverUCase: driverUCase}

	g.GET("/drivers/:id", httpCommon.Auth(fAuth), h.getDriver)
	g.POST("/drivers", httpCommon.Auth(fAuth), h.addDriver)
	g.PUT("/drivers/:id", httpCommon.Auth(fAuth), h.updateDriver)

	return h
}
