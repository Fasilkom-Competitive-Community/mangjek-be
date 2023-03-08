package http

import "time"

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
)
