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
		ID    string
		Email string

		Name        string
		NIM         string
		PhoneNumber string

		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddUser struct {
		ID    string
		Email string

		Name        string
		NIM         string
		PhoneNumber string
	}
	UpdateUser struct {
		ID    string
		Email string

		Name        string
		NIM         string
		PhoneNumber string
	}
)
