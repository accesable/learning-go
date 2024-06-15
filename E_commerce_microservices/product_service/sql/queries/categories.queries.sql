-- name: CreateCategory :execresult
INSERT INTO categories (name)
VALUES (?);
-- name: GetCategoryByID :one
SELECT * 
FROM categories
WHERE id = ?;
-- name: GetCategoryByName :one
SELECT * 
FROM categories
WHERE name = ?;
-- name: UpdateCategory :execresult
UPDATE categories
SET name = ?, updated_at = CURRENT_TIMESTAMP
WHERE id = ?;
-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = ?;
-- name: ListCategories :many
SELECT * 
FROM categories
ORDER BY created_at DESC;

