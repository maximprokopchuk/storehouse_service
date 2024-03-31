// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package sqlc

import (
	"context"
)

const addItemForStorehouse = `-- name: AddItemForStorehouse :one
INSERT INTO items (
  detail_id, storehouse_id
) VALUES (
  $1, $2
)
RETURNING id, storehouse_id, detail_id, count
`

type AddItemForStorehouseParams struct {
	DetailID     int32
	StorehouseID int32
}

func (q *Queries) AddItemForStorehouse(ctx context.Context, arg AddItemForStorehouseParams) (Item, error) {
	row := q.db.QueryRow(ctx, addItemForStorehouse, arg.DetailID, arg.StorehouseID)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.StorehouseID,
		&i.DetailID,
		&i.Count,
	)
	return i, err
}

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

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
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

const getAllItemsInStorehouse = `-- name: GetAllItemsInStorehouse :many
SELECT id, storehouse_id, detail_id, count FROM items
WHERE storehouse_id = $1
`

func (q *Queries) GetAllItemsInStorehouse(ctx context.Context, storehouseID int32) ([]Item, error) {
	rows, err := q.db.Query(ctx, getAllItemsInStorehouse, storehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.StorehouseID,
			&i.DetailID,
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

const getStorehousesForCity = `-- name: GetStorehousesForCity :many
SELECT id, city_id, name FROM storehouse
WHERE city_id = $1
`

func (q *Queries) GetStorehousesForCity(ctx context.Context, cityID int32) ([]Storehouse, error) {
	rows, err := q.db.Query(ctx, getStorehousesForCity, cityID)
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

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET count = $2
WHERE id = $1
RETURNING id, storehouse_id, detail_id, count
`

type UpdateItemParams struct {
	ID    int64
	Count int32
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRow(ctx, updateItem, arg.ID, arg.Count)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.StorehouseID,
		&i.DetailID,
		&i.Count,
	)
	return i, err
}