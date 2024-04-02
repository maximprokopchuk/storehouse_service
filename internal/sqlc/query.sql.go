// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
)

const createStorehouse = `-- name: CreateStorehouse :one
INSERT INTO storehouse (
  name, city_id
) VALUES (
  $1, $2
)
RETURNING id, city_id, name
`

type CreateStorehouseParams struct {
	Name   string
	CityID int32
}

func (q *Queries) CreateStorehouse(ctx context.Context, arg CreateStorehouseParams) (Storehouse, error) {
	row := q.db.QueryRow(ctx, createStorehouse, arg.Name, arg.CityID)
	var i Storehouse
	err := row.Scan(&i.ID, &i.CityID, &i.Name)
	return i, err
}

const createStorehouseItemForStorehouse = `-- name: CreateStorehouseItemForStorehouse :one
INSERT INTO storehouse_item (
  component_id, storehouse_id
) VALUES (
  $1, $2
)
RETURNING id, storehouse_id, component_id, count
`

type CreateStorehouseItemForStorehouseParams struct {
	ComponentID  int32
	StorehouseID int32
}

func (q *Queries) CreateStorehouseItemForStorehouse(ctx context.Context, arg CreateStorehouseItemForStorehouseParams) (StorehouseItem, error) {
	row := q.db.QueryRow(ctx, createStorehouseItemForStorehouse, arg.ComponentID, arg.StorehouseID)
	var i StorehouseItem
	err := row.Scan(
		&i.ID,
		&i.StorehouseID,
		&i.ComponentID,
		&i.Count,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM storehouse_item
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteItem, id)
	return err
}

const deleteStorehouse = `-- name: DeleteStorehouse :exec
DELETE FROM storehouse
WHERE id = $1
`

func (q *Queries) DeleteStorehouse(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteStorehouse, id)
	return err
}

const getAllStorehouseItemsByStorehouse = `-- name: GetAllStorehouseItemsByStorehouse :many
SELECT id, storehouse_id, component_id, count FROM storehouse_item
WHERE storehouse_id = $1
`

func (q *Queries) GetAllStorehouseItemsByStorehouse(ctx context.Context, storehouseID int32) ([]StorehouseItem, error) {
	rows, err := q.db.Query(ctx, getAllStorehouseItemsByStorehouse, storehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StorehouseItem
	for rows.Next() {
		var i StorehouseItem
		if err := rows.Scan(
			&i.ID,
			&i.StorehouseID,
			&i.ComponentID,
			&i.Count,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getStorehousesByCity = `-- name: GetStorehousesByCity :many
SELECT id, city_id, name FROM storehouse
WHERE city_id = $1
`

func (q *Queries) GetStorehousesByCity(ctx context.Context, cityID int32) ([]Storehouse, error) {
	rows, err := q.db.Query(ctx, getStorehousesByCity, cityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Storehouse
	for rows.Next() {
		var i Storehouse
		if err := rows.Scan(&i.ID, &i.CityID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStorehouseItem = `-- name: UpdateStorehouseItem :one
UPDATE storehouse_item
SET count = $2
WHERE id = $1
RETURNING id, storehouse_id, component_id, count
`

type UpdateStorehouseItemParams struct {
	ID    int64
	Count int32
}

func (q *Queries) UpdateStorehouseItem(ctx context.Context, arg UpdateStorehouseItemParams) (StorehouseItem, error) {
	row := q.db.QueryRow(ctx, updateStorehouseItem, arg.ID, arg.Count)
	var i StorehouseItem
	err := row.Scan(
		&i.ID,
		&i.StorehouseID,
		&i.ComponentID,
		&i.Count,
	)
	return i, err
}
