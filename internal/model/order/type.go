package order

import (
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
	uModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/user"

	"time"
)

const (
	OnProgressStatus Status = "ON PROGRESS"
	PickedUpStatus   Status = "PICKED UP"
	ArrivedStatus    Status = "ARRIVED"
)

type (
	Status string

	OrderInquiry struct {
		ID           string
		UserID       string
		OrderRouteID int32
		Price        int64
		Distance     int32
		Duration     int32
		Origin       Location
		Destination  Location
		Routes       string

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddOrderInquiry struct {
		ID          string
		UserID      string
		Price       int64
		Distance    int32
		Duration    int32
		Origin      Location
		Destination Location
		Routes      string
	}

	Order struct {
		ID           string
		User         uModel.User
		DName        string
		Driver       dModel.Driver
		OrderInquiry OrderInquiry
		Payment      pModel.Payment
		Status       Status

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddOrder struct {
		ID             string
		UserID         string
		DriverID       int32
		OrderInquiryID string
		Status         Status

		Payment pModel.AddPayment
	}

	Location struct {
		Address   string
		Latitude  float64
		Longitude float64
	}

	Direction struct {
		Distance         int32
		Duration         int32
		Origin           Location
		Destination      Location
		OverviewPolyline []Location
	}
)
