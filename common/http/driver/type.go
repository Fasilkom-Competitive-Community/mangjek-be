package http

import "time"

type (
	Driver struct {
		ID           int32     `json:"id"`
		UserID       string    `json:"user_id"`
		PoliceNumber string    `json:"police_number"`
		VehicleModel string    `json:"vehicle_model"`
		VehicleType  string    `json:"vehicle_type"`
		Nik          string    `json:"nik"`
		Address      string    `json:"address"`
		IsSimActive  bool      `json:"is_sim_active"`
		IsStnkActive bool      `json:"is_stnk_active"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
	AddDriver struct {
		UserID       string `json:"user_id" validate:"required"`
		PoliceNumber string `json:"police_number" validate:"required"`
		VehicleModel string `json:"vehicle_model" validate:"required"`
		VehicleType  string `json:"vehicle_type" validate:"required,oneof=MOTOR MOBIL"`
		Nik          string `json:"nik" validate:"required"`
		Address      string `json:"address" validate:"required"`
		IsSimActive  bool   `json:"is_sim_active" validate:"required"`
		IsStnkActive bool   `json:"is_stnk_active" validate:"required"`
	}
)
