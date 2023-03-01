package http

import (
	"github.com/gin-gonic/gin"
	"net/http"

	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (d HTTPUserDelivery) addUser(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	var user httpCommon.AddUser
	if err := c.ShouldBindJSON(&user); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.userUCase.CreateUser(ctx, uModel.AddUser{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: nid,
	})
}
