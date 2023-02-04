package user

import (
	"time"
)

type (
	AuthDriver struct {
		ID   string
		Role uint32
	}
	AuthDriverFull struct {
		IsEmailVerified bool
	}
	Driver struct {
		ID     int32
		UserID string
		Email  string

		Nama        string
		NIK         string
		PhoneNumber string

		Alamat     string
		statusStnk bool
		statusSim  bool

		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddDriver struct {
		ID     int32
		UserID string
		Email  string

		Nama        string
		NIK         string
		PhoneNumber string

		Alamat     string
		statusStnk bool
		statusSim  bool
	}
	UpdateDriver struct {
		ID     int32
		UserID string
		Email  string

		Nama        string
		NIK         string
		PhoneNumber string

		Alamat     string
		statusStnk bool
		statusSim  bool
	}
)
