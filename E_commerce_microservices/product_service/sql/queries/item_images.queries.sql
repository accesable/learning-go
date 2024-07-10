-- name: ListAllImages :many
SELECT * FROM item_images;

-- name: ListImagesByItemId :many
SELECT * FROM item_images WHERE item_id = ?;

-- name: InsertImages :execresult 
INSERT INTO item_images (
  display_name, image_url, item_id
) VALUES ( ?,?,? )

DELETE FROM item_images
  WHERE id = ?;
