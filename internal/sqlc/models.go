// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package sqlc

type Item struct {
	ID           int64
	StorehouseID int32
	DetailID     int32
	Count        int32
}

type Storehouse struct {
	ID     int64
	CityID int32
	Name   string
}
