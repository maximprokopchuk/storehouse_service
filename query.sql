-- name: CreateStorehouse :one
INSERT INTO storehouse (
  name, city_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetStorehousesByCity :many
SELECT * FROM storehouse
WHERE city_id = $1;


-- name: GetAllStorehouseItemsByStorehouse :many
SELECT * FROM storehouse_item
WHERE storehouse_id = $1;

-- name: CreateStorehouseItemForStorehouse :one
INSERT INTO storehouse_item (
  component_id, storehouse_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateStorehouseItem :one
UPDATE storehouse_item
SET count = $2
WHERE id = $1
RETURNING *;


-- name: DeleteStorehouse :exec
DELETE FROM storehouse
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM storehouse_item
WHERE id = $1;