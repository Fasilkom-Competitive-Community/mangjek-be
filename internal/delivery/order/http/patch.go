package http

import (
	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	httpOrderCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/order"
	oModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/order"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPOrderDelivery) updateOrderStatus(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	var uos httpOrderCommon.UpdateOrderStatus
	if err := c.ShouldBindQuery(&uos); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	err := d.orderUCase.UpdateOrderStatus(ctx, oModel.UpdateOrderStatus{
		ID:     id,
		Status: oModel.Status(uos.Status),
	}, au)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Data: id,
	})
}
