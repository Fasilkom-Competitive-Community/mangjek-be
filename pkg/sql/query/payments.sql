-- name: GetPayment :one
SELECT id
     , amount
     , status
     , method
     , qr_str
     , created_at
     , updated_at
FROM payments
WHERE id = $1;

-- name: CreatePayment :one
INSERT INTO payments ( id
                     , amount
                     , status
                     , method
                     , qr_str)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: UpdatePaymentStatusToPaid :one
UPDATE payments
SET status     = $2
  , updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;