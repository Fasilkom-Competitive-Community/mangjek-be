-- name: ListUsers :many
SELECT id
     , name
     , email
     , phone_number
     , nim
     , created_at
     , updated_at
FROM users;

-- name: GetUser :one
SELECT id
     , name
     , email
     , phone_number
     , nim
     , created_at
     , updated_at
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users ( id
                  , name
                  , email
                  , phone_number
                  , nim)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdateUser :one
UPDATE users
SET name         = $2
  , email        = $3
  , phone_number = $4
  , nim          = $5
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;