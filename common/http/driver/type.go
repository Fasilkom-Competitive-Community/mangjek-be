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
		UserID       string `json:"user_id" binding:"required"`
		PoliceNumber string `json:"police_number" binding:"required"`
		VehicleModel string `json:"vehicle_model" binding:"required"`
		VehicleType  string `json:"vehicle_type" binding:"required,oneof=MOTOR MOBIL"`
		Nik          string `json:"nik" binding:"required"`
		Address      string `json:"address" binding:"required"`
		IsSimActive  bool   `json:"is_sim_active" binding:"required"`
		IsStnkActive bool   `json:"is_stnk_active" binding:"required"`
	}

	UpdateDriver struct {
		UserID       string `json:"user_id" binding:"required"`
		PoliceNumber string `json:"police_number" binding:"required"`
		VehicleModel string `json:"vehicle_model" binding:"required"`
		VehicleType  string `json:"vehicle_type" binding:"required,oneof=MOTOR MOBIL"`
		Nik          string `json:"nik" binding:"required"`
		Address      string `json:"address" binding:"required"`
		IsSimActive  bool   `json:"is_sim_active" binding:"required"`
		IsStnkActive bool   `json:"is_stnk_active" binding:"required"`
	}
)
