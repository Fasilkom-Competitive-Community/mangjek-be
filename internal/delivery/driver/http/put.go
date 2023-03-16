package http

import (
	"fmt"
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	httpDriverCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/driver"
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPDriverDelivery) updateDriver(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	fmt.Println("Aman")

	//id, err := strconv.Atoi(c.Param("id"))
	//if err != nil {
	//	panic(err)
	//}

	var driver httpDriverCommon.UpdateDriver
	if err := c.ShouldBindJSON(&driver); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return

	}

	fmt.Println("Aman1")

	nid, err := d.driverUCase.UpdateDriver(ctx, dModel.UpdateDriver{
		UserID:       au.ID,
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

	fmt.Println(nid)

	c.JSON(http.StatusOK, httpCommon.Response{
		Data: nid,
	})
}
