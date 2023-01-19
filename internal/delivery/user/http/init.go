package http

import (
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	uUCase "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type HTTPUserDelivery struct {
	userUCase uUCase.Usecase
}

func NewHTTPUserDelivery(g *gin.RouterGroup, userUCase uUCase.Usecase, fAuth *auth.Client) HTTPUserDelivery {
	h := HTTPUserDelivery{userUCase: userUCase}

	g.GET("/users/:id", h.getUser, httpCommon.Auth(fAuth))
	g.POST("/users", h.addUser, httpCommon.Auth(fAuth))
	g.PUT("/users/:id", h.updateUser, httpCommon.Auth(fAuth))

	return h
}
