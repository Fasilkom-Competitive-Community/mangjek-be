-- name: ListDrivers :many
SELECT id
     , user_id
     , police_number
     , vehicle_model
     , vehicle_type
     , nik
     , address
     , is_sim_active
     , is_stnk_active
     , created_at
     , updated_at
FROM drivers;

-- name: GetDriver :one
SELECT id
     , user_id
     , police_number
     , vehicle_model
     , vehicle_type
     , nik
     , address
     , is_sim_active
     , is_stnk_active
     , created_at
     , updated_at
FROM drivers
WHERE id = $1;

-- name: GetDriverByUserID :one
SELECT id
     , user_id
     , police_number
     , vehicle_model
     , vehicle_type
     , nik
     , address
     , is_sim_active
     , is_stnk_active
     , created_at
     , updated_at
FROM drivers
WHERE user_id = $1;

-- name: CreateDriver :one
INSERT INTO drivers ( user_id
                    , police_number
                    , vehicle_model
                    , vehicle_type
                    , nik
                    , address
                    , is_sim_active
                    , is_stnk_active)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: UpdateDriver :one
UPDATE drivers
SET police_number  = $2
  , vehicle_model  = $3
  , vehicle_type   = $4
  , nik            = $5
  , address        = $6
  , is_sim_active  = $7
  , is_stnk_active = $8
WHERE user_id = $1
RETURNING id;

-- name: DeleteDriver :exec
DELETE
FROM drivers
WHERE id = $1;