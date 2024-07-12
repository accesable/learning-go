-- name: ListProducts :many
SELECT * FROM items;
-- name: ListProductsWithCategory :many
SELECT 
    i.id, 
    i.name, 
    i.original_price, 
    i.short_description, 
    c.name AS category_name
FROM 
    items i
JOIN 
    categories c 
ON 
    i.category_id = c.id;
-- name: InsertProduct :execresult
INSERT INTO items (
  name,original_price,short_description,category_id
) VALUES ( ?,?,?,? );
-- name: DeleteProduct :execresult
DELETE FROM items
  WHERE id = ?;
-- name: UpdateProduct :execresult
UPDATE items
  SET name = ? , original_price = ? , short_description = ? , updated_at = CURRENT_TIMESTAMP
  WHERE condition;

-- name: PartialUpdateItem :execresult
UPDATE items
SET
    name = COALESCE(?, name),
    category_id = COALESCE(?, category_id),
    short_description = COALESCE(?, short_description),
    original_price = COALESCE(?, original_price),
    updated_at = CURRENT_TIMESTAMP 
WHERE id = ?
