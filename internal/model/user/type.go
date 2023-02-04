package user

import (
	"time"
)

type (
	AuthUser struct {
		ID   string
		Role uint32
	}
	AuthUserFull struct {
		IsEmailVerified bool
	}
	User struct {
		ID          string
		Name        string
		Email       string
		PhoneNumber string
		Nim         string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddUser struct {
		ID          string
		Name        string
		Email       string
		PhoneNumber string
		Nim         string
	}
	UpdateUser struct {
		ID          string
		Name        string
		Email       string
		PhoneNumber string
		Nim         string
	}
)
