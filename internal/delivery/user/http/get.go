package http

import (
	httpOrderCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/order"
	"github.com/gin-gonic/gin"
	"net/http"

	httpCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
)

func (d HTTPUserDelivery) getUser(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	u, err := d.userUCase.GetUser(ctx, id, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: httpCommon.User(u)})
}

func (d HTTPUserDelivery) getUserHistory(c *gin.Context) {
	ctx := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	oh, err := d.userUCase.GetUserHistory(ctx, id, au)
	if err != nil {
		c.Error(err)
		return
	}

	var resp []httpOrderCommon.GetOrder
	for i := 0; i < len(oh); i++ {
		resp = append(resp, httpOrderCommon.GetOrder{
			ID:    oh[i].ID,
			UName: oh[i].User.Name,
			Driver: httpOrderCommon.Driver{
				Name:         oh[i].DName,
				PoliceNumber: oh[i].Driver.PoliceNumber,
				VehicleType:  oh[i].Driver.VehicleType,
				VehicleModel: oh[i].Driver.VehicleModel,
			},
			OrderInquiry: httpOrderCommon.Inquiry{
				Price:    oh[i].OrderInquiry.Price,
				Distance: oh[i].OrderInquiry.Distance,
				Duration: oh[i].OrderInquiry.Duration,
				OAddress: oh[i].OrderInquiry.Origin.Address,
				DAddress: oh[i].OrderInquiry.Destination.Address,
			},
			Payment: httpOrderCommon.Payment{
				Amount:   oh[i].Payment.Amount,
				Status:   string(oh[i].Payment.Status),
				Method:   string(oh[i].Payment.Method),
				QrString: oh[i].Payment.QrString,
			},
			Status:    string(oh[i].Status),
			CreatedAt: oh[i].CreatedAt,
			UpdatedAt: oh[i].UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: resp})
}
