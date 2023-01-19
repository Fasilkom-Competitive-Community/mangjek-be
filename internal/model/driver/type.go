package user

import (
	"time"
)

type (
	Driver struct {
		ID     int32
		UserID string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddDriver struct {
		ID     int32
		UserID string
	}
	UpdateDriver struct {
		ID     int32
		UserID string
	}
)
