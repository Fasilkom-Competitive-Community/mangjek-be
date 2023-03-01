package order

import "time"

type (
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
