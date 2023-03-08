package order

import (
	dModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/driver"
	pModel "github.com/Fasilkom-Competitive-Community/mangjek-be/internal/model/payment"
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
		ID      string
		Driver  dModel.Driver
		Payment pModel.Payment
	}

	AddOrder struct {
		ID             string
		UserID         string
		DriverID       string
		OrderInquiryID string
		PaymentID      string
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
