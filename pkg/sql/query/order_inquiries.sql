-- name: GetOrderInquiry :one
SELECT id
     , user_id
     , price
     , distance
     , duration
     , origin_lat
     , origin_long
     , origin_address
     , destination_lat
     , destination_long
     , destination_address
     , routes
     , created_at
     , updated_at
FROM order_inquiries
WHERE id = $1;

-- name: CreateOrderInquiry :one
INSERT INTO order_inquiries ( id
                            , user_id
                            , price
                            , distance
                            , duration
                            , origin_lat
                            , origin_long
                            , origin_address
                            , destination_lat
                            , destination_long
                            , destination_address
                            , routes)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id;

-- name: DeleteOrderInquiry :exec
DELETE
FROM order_inquiries
WHERE id = $1;