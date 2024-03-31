-- name: CreateStorehouse :one
INSERT INTO storehouse (
  name, city_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetStorehousesForCity :many
SELECT * FROM storehouse
WHERE city_id = $1;


-- name: GetAllItemsInStorehouse :many
SELECT * FROM items
WHERE storehouse_id = $1;

-- name: AddItemForStorehouse :one
INSERT INTO items (
  detail_id, storehouse_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateItem :one
UPDATE items
SET count = $2
WHERE id = $1
RETURNING *;


-- name: DeleteStorehouse :exec
DELETE FROM storehouse
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;