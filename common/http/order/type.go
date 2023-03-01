package http

import "time"

type (
	Location struct {
		Address   string  `json:"address"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	AddLocation struct {
		Latitude  float64 `json:"latitude" validate:"required,latitude"`
		Longitude float64 `json:"longitude" validate:"required,longitude"`
	}

	OrderInquiry struct {
		ID          string    `json:"id"`
		UserID      string    `json:"user_id"`
		Price       int64     `json:"price"`
		Distance    int32     `json:"distance"`
		Duration    int32     `json:"duration"`
		Origin      Location  `json:"origin"`
		Destination Location  `json:"destination"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	AddOrderInquiry struct {
		UserID      string      `json:"user_id" validate:"required"`
		Origin      AddLocation `json:"origin"`
		Destination AddLocation `json:"destination"`
	}
)
