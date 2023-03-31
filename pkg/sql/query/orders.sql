-- name: GetOrder :one
SELECT *
FROM orders
         JOIN users ON orders.user_id = users.id
         JOIN drivers ON orders.driver_id = drivers.id
         JOIN order_inquiries ON orders.order_inquiry_id = order_inquiries.id
         JOIN payments ON orders.payment_id = payments.id
WHERE orders.id = $1;

-- name: CreateOrder :one
INSERT INTO orders ( id
                   , user_id
                   , driver_id
                   , order_inquiry_id
                   , payment_id
                   , status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
