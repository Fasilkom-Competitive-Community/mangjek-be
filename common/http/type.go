package http

import "time"

type (
	Error struct {
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Response struct {
		Data any `json:"data"`
	}

	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	AddUser struct {
		ID          string `json:"id" validate:"required"`
		Email       string `json:"email" validate:"required,email,contains=@student.unsri.ac.id"`
		Name        string `json:"name" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
	}
	UpdateUser struct {
		ID          string `json:"id" validate:"required"`
		Email       string `json:"email" validate:"required,email,contains=@student.unsri.ac.id"`
		Name        string `json:"name" validate:"required"`
		PhoneNumber string `json:"phone_number" validate:"required"`
	}
)
