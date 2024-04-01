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


-- name: GetAllItemsByStorehouse :many
SELECT * FROM item
WHERE storehouse_id = $1;

-- name: CreateItemForStorehouse :one
INSERT INTO item (
  component_id, storehouse_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateItem :one
UPDATE item
SET count = $2
WHERE id = $1
RETURNING *;


-- name: DeleteStorehouse :exec
DELETE FROM storehouse
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM item
WHERE id = $1;