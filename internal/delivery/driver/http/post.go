package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"

	httpDriverCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/driver"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
)

func (d HTTPDriverDelivery) addDriver(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	var driver httpDriverCommon.AddDriver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.driverUCase.CreateDriver(ctx, dModel.AddDriver{
		UserID:       driver.UserID,
		PoliceNumber: driver.PoliceNumber,
		VehicleModel: driver.VehicleModel,
		VehicleType:  driver.VehicleType,
		Nik:          driver.Nik,
		Address:      driver.Address,
		IsSimActive:  driver.IsSimActive,
		IsStnkActive: driver.IsStnkActive,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: nid,
	})
}
