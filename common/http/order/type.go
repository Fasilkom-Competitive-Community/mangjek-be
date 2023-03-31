package http

import (
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

	Driver struct {
		Name         string `json:"name"`
		PoliceNumber string `json:"police_number"`
		VehicleModel string `json:"vehicle_model"`
		VehicleType  string `json:"vehicle_type"`
	}

	Inquiry struct {
		Price    int64  `json:"price"`
		Distance int32  `json:"distance"`
		Duration int32  `json:"duration"`
		OAddress string `json:"o_address"`
		DAddress string `json:"d_address"`
	}

	Payment struct {
		Amount   float64 `json:"amount"`
		Status   string  `json:"status"`
		Method   string  `json:"method"`
		QrString string  `json:"qr_string"`
	}

	GetOrder struct {
		ID           string    `json:"id"`
		UName        string    `json:"u_name"`
		Driver       Driver    `json:"driver"`
		OrderInquiry Inquiry   `json:"order_inquiry"`
		Payment      Payment   `json:"payment"`
		Status       string    `json:"status"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	GetOrderRoutes struct {
		ID     string     `json:"id"`
		Routes []Location `json:"routes"`
	}
)
