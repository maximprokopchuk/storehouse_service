package grpcserver

import (
	"context"

	"github.com/maximprokopchuk/storehouse_service/internal/sqlc"
	"github.com/maximprokopchuk/storehouse_service/internal/store"
	"github.com/maximprokopchuk/storehouse_service/pkg/api"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateStorehouse(ctx context.Context, req *api.CreateStorehouseRequest) (*api.Storehouse, error) {
	params := sqlc.CreateStorehouseParams{
		Name:   req.GetName(),
		CityID: req.GetCityId(),
	}
	rec, err := server.Store.Queries.CreateStorehouse(ctx, params)

	if err != nil {
		return nil, err
	}
	return &api.Storehouse{Id: int32(rec.ID), Name: rec.Name, CityId: rec.CityID}, nil
}

func (server *GRPCServer) GetStorehousesListByCityId(ctx context.Context, req *api.GetStorehousesListByCityIdRequest) (*api.GetStorehousesListResponse, error) {
	rec, err := server.Store.Queries.GetStorehousesByCity(ctx, req.GetCityId())

	if err != nil {
		return nil, err
	}

	var storehouses []*api.Storehouse

	for _, storehouse := range rec {
		newRec := api.Storehouse{
			CityId: storehouse.CityID,
			Name:   storehouse.Name,
			Id:     int32(storehouse.ID),
		}
		storehouses = append(storehouses, &newRec)
	}

	return &api.GetStorehousesListResponse{Items: storehouses}, nil
}

func (server *GRPCServer) GetItemsByStorehouseId(ctx context.Context, req *api.GetItemsByStorehouseIdRequest) (*api.GetItemsByStorehouseIdResponse, error) {
	rec, err := server.Store.Queries.GetAllItemsByStorehouse(ctx, req.GetStorehouseId())

	if err != nil {
		return nil, err
	}

	var items []*api.Item

	for _, item := range rec {
		newRec := api.Item{
			Id:    int32(item.ID),
			Count: item.Count,
		}
		items = append(items, &newRec)
	}

	return &api.GetItemsByStorehouseIdResponse{Items: items}, nil
}

func (server *GRPCServer) CreateItemForStorehouse(ctx context.Context, req *api.CreateItemForStorehoseRequest) (*api.Item, error) {
	params := sqlc.CreateItemForStorehouseParams{
		ComponentID:  req.GetComponentId(),
		StorehouseID: req.GetStorehouseId(),
	}
	rec, err := server.Store.Queries.CreateItemForStorehouse(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.Item{
		Id:           int32(rec.ID),
		ComponentId:  rec.ComponentID,
		StorehouseId: rec.StorehouseID,
		Count:        rec.Count,
	}, nil
}

func (server *GRPCServer) UpdateItem(ctx context.Context, req *api.UpdateItemRequest) (*api.Item, error) {
	params := sqlc.UpdateItemParams{
		ID:    int64(req.GetId()),
		Count: req.GetCount(),
	}
	rec, err := server.Store.Queries.UpdateItem(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.Item{
		Id:           int32(rec.ID),
		ComponentId:  rec.ComponentID,
		StorehouseId: rec.StorehouseID,
		Count:        rec.Count,
	}, nil
}

func (server *GRPCServer) DeleteStorehouse(ctx context.Context, req *api.DeleteStorehouseRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteStorehouse(ctx, int64(req.GetId()))

	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}

func (server *GRPCServer) DeleteItem(ctx context.Context, req *api.DeleteItemRequest) (*api.Empty, error) {
	err := server.Store.Queries.DeleteItem(ctx, int64(req.GetId()))

	if err != nil {
		return nil, err
	}

	return &api.Empty{}, nil
}
