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

	g.GET("/users/:id", httpCommon.Auth(fAuth), h.getUser)
	g.POST("/users", httpCommon.Auth(fAuth), h.addUser)
	g.PUT("/users/:id", httpCommon.Auth(fAuth), h.updateUser)

	return h
}
