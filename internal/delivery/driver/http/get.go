package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	httpDriverCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/driver"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (d HTTPDriverDelivery) getDriver(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	id := c.Param("id")
	idn, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Error(err)
		return
	}

	u, err := d.driverUCase.GetDriver(ctx, int32(idn), au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: httpDriverCommon.Driver(u)})
}
