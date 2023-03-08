package http

import (
	gcs "cloud.google.com/go/storage"
	"firebase.google.com/go/v4/auth"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	"github.com/gin-gonic/gin"
)

type HTTPFileDelivery struct {
	//uUCase.Usecase
	bkt *gcs.BucketHandle
}

func NewHTTPFileDelivery(g *gin.RouterGroup, fAuth *auth.Client) HTTPFileDelivery {
	h := HTTPFileDelivery{}

	//g.GET("/file", httpCommon.Auth(fAuth), h.getImage)
	g.POST("/file", httpCommon.Auth(fAuth), h.uploadProfile)
	//g.PUT("/drivers", httpCommon.Auth(fAuth), h.updateDriver)

	return h
}
