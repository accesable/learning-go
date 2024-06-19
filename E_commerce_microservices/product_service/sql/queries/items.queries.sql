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

