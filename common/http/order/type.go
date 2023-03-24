package http

import (
	"github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	"github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"
	"google.golang.org/grpc/status"
	"time"

	httpPaymentCommon "github.com/Fasilkom-Competitive-Community/mangjek-be/common/http/payment"
)

type (
	Location struct {
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	AddLocation struct {
		Latitude  float64 `json:"latitude" binding:"required,latitude"`
		Longitude float64 `json:"longitude" binding:"required,longitude"`
	}

	OrderInquiry struct {
		ID          string     `json:"id"`
		UserID      string     `json:"user_id"`
		Price       int64      `json:"price"`
		Distance    int32      `json:"distance"`
		Duration    int32      `json:"duration"`
		Origin      Location   `json:"origin"`
		Destination Location   `json:"destination"`
		Routes      []Location `json:"routes"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}

	AddOrderInquiry struct {
		UserID      string   `json:"user_id" binding:"required"`
		Origin      Location `json:"origin"`
		Destination Location `json:"destination"`
	}

	AddOrder struct {
		DriverID       int32  `json:"driver_id" binding:"required"`
		UserID         string `json:"user_id" binding:"required"`
		OrderInquiryID string `json:"order_inquiry_id" binding:"required"`

		httpPaymentCommon.AddPayment
	}

	Order struct {
		ID          string     `json:"id"`
		UserID      string     `json:"user_id"`
		Price       int64      `json:"price"`
		Distance    int32      `json:"distance"`
		Duration    int32      `json:"duration"`
		Origin      Location   `json:"origin"`
		Destination Location   `json:"destination"`
		Routes      []Location `json:"routes"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
	}

	GetOrder struct {
		ID           string          `json:"id"`
		User         user.User       `json:"user"`
		DName        string          `json:"d_name"`
		Driver       driver.Driver   `json:"driver"`
		OrderInquiry OrderInquiry    `json:"order_inquiry"`
		Payment      payment.Payment `json:"payment"`
		Status       status.Status   `json:"status"`
		CreatedAt    time.Time       `json:"created_at"`
		UpdatedAt    time.Time       `json:"updated_at"`
	}
)
