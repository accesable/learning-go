-- name: CreateProduct :exec
INSERT INTO products (name, description, image, price, quantity)
VALUES (?, ?, ?, ?, ?);

-- name: GetProductByID :one
SELECT id, name, description, image, price, quantity, createdAt
FROM products
WHERE id = ?;

-- name: UpdateProduct :exec
UPDATE products
SET name = ?, description = ?, image = ?, price = ?, quantity = ?
WHERE id = ?;

-- name: DeleteProduct :exec
DELETE FROM products
WHERE id = ?;

-- name: ListProducts :many
SELECT id, name, description, image, price, quantity, createdAt
FROM products
ORDER BY createdAt DESC;
