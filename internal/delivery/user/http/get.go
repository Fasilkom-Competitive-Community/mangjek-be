package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (d HTTPUserDelivery) getUser(c *gin.Context) {
	ctx := c.Request.Context()
	auStr, _ := c.Get(httpCommon.AUTH_USER)
	au := auStr.(uModel.AuthUser)

	id := c.Param("id")

	u, err := d.userUCase.GetUser(ctx, id, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: httpCommon.User(u)})
}
