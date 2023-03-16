package driver

import (
	"time"
)

type (
	Driver struct {
		ID           int32
		UserID       string
		PoliceNumber string
		VehicleModel string
		VehicleType  string
		Nik          string
		Address      string
		IsSimActive  bool
		IsStnkActive bool

		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddDriver struct {
		UserID       string
		PoliceNumber string
		VehicleModel string
		VehicleType  string
		Nik          string
		Address      string
		IsSimActive  bool
		IsStnkActive bool
	}
	UpdateDriver struct {
		UserID       string
		PoliceNumber string
		VehicleModel string
		VehicleType  string
		Nik          string
		Address      string
		IsSimActive  bool
		IsStnkActive bool
	}
	getDriverName struct {
		UserID string
		Name   string
	}
)
