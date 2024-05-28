-- name: ListProducts :many
SELECT * FROM `products`
ORDER BY name ;
-- name: GetProduct :one
SELECT * FROM `products`
WHERE id = ? LIMIT 1;
-- name: ListCategories :many
SELECT * FROM `categories`
ORDER BY name ;
-- name: GetCategory :one
SELECT * FROM `categories`
WHERE id = ? LIMIT 1;