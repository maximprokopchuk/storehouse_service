package grpcserver

import (
	"context"

	"github.com/maximprokopchuk/storehouse_service/internal/sqlc"
	"github.com/maximprokopchuk/storehouse_service/internal/store"
	"github.com/maximprokopchuk/storehouse_service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	Store *store.Store
}

func New(st *store.Store) *GRPCServer {
	return &GRPCServer{Store: st}
}

func (server *GRPCServer) CreateStorehouse(ctx context.Context, req *api.CreateStorehouseRequest) (*api.CreateStorehouseResponse, error) {
	params := sqlc.CreateStorehouseParams{
		Name:   req.GetName(),
		CityID: req.GetCityId(),
	}
	rec, err := server.Store.Queries.CreateStorehouse(ctx, params)

	if err != nil {
		return nil, err
	}
	return &api.CreateStorehouseResponse{
		Result: &api.Storehouse{
			Id:     int32(rec.ID),
			Name:   rec.Name,
			CityId: rec.CityID,
		},
	}, nil
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

	return &api.GetStorehousesListResponse{Result: storehouses}, nil
}

func (server *GRPCServer) GetStorehouseItemById(ctx context.Context, req *api.GetStorehouseItemByIdRequest) (*api.GetStorehouseItemByIdResponse, error) {
	rec, err := server.Store.Queries.GetAllStorehouseItemById(ctx, int64(req.GetId()))

	if err != nil {
		return nil, err
	}

	return &api.GetStorehouseItemByIdResponse{
		Result: &api.StorehouseItem{
			Id:           int32(rec.ID),
			Count:        rec.Count,
			StorehouseId: rec.StorehouseID,
			ComponentId:  rec.ComponentID,
		},
	}, nil
}

func (server *GRPCServer) GetStorehouseItemsByStorehouseId(ctx context.Context, req *api.GetStorehouseItemsByStorehouseIdRequest) (*api.GetStorehouseItemsByStorehouseIdResponse, error) {
	rec, err := server.Store.Queries.GetAllStorehouseItemsByStorehouse(ctx, req.GetStorehouseId())

	if err != nil {
		return nil, err
	}

	var items []*api.StorehouseItem

	for _, item := range rec {
		newRec := api.StorehouseItem{
			Id:           int32(item.ID),
			Count:        item.Count,
			ComponentId:  item.ComponentID,
			StorehouseId: item.StorehouseID,
		}
		items = append(items, &newRec)
	}

	return &api.GetStorehouseItemsByStorehouseIdResponse{Result: items}, nil
}

func (server *GRPCServer) GetStorehouseItemsByStorehouseIdAndComponentsIds(ctx context.Context, req *api.GetStorehouseItemsByStorehouseIdAndComponentsIdsRequest) (*api.GetStorehouseItemsByStorehouseIdAndComponentsIdsResponse, error) {
	params := sqlc.GetStorehouseItemsByStorehouseAndComponentsParams{
		StorehouseID:  req.GetStorehouseId(),
		ComponentsIds: req.GetComponentIds(),
	}
	rec, err := server.Store.Queries.GetStorehouseItemsByStorehouseAndComponents(ctx, params)

	if err != nil {
		return nil, err
	}

	var items []*api.StorehouseItem

	for _, item := range rec {
		newRec := api.StorehouseItem{
			Id:           int32(item.ID),
			Count:        item.Count,
			ComponentId:  item.ComponentID,
			StorehouseId: item.StorehouseID,
		}
		items = append(items, &newRec)
	}

	return &api.GetStorehouseItemsByStorehouseIdAndComponentsIdsResponse{Result: items}, nil
}

func (server *GRPCServer) CreateStorehouseItemForStorehouse(ctx context.Context, req *api.CreateStorehouseItemForStorehoseRequest) (*api.CreateStorehouseItemResponse, error) {
	existing_params := sqlc.GetStorehouseItemsByStorehouseAndComponentIdParams{
		ComponentID:  req.GetComponentId(),
		StorehouseID: req.GetStorehouseId(),
	}

	existing, err := server.Store.Queries.GetStorehouseItemsByStorehouseAndComponentId(ctx, existing_params)

	if err != nil {
		return nil, err
	}

	if len(existing) != 0 {
		err := status.Error(codes.AlreadyExists, "item for component already exists")
		return nil, err
	}

	params := sqlc.CreateStorehouseItemForStorehouseParams{
		ComponentID:  req.GetComponentId(),
		StorehouseID: req.GetStorehouseId(),
		Count:        req.GetCount(),
	}

	rec, err := server.Store.Queries.CreateStorehouseItemForStorehouse(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.CreateStorehouseItemResponse{
		Result: &api.StorehouseItem{
			Id:           int32(rec.ID),
			ComponentId:  rec.ComponentID,
			StorehouseId: rec.StorehouseID,
			Count:        rec.Count,
		},
	}, nil
}

func (server *GRPCServer) UpdateStorehouseItem(ctx context.Context, req *api.UpdateStorehouseItemRequest) (*api.UpdateStorehouseItemResponse, error) {
	params := sqlc.UpdateStorehouseItemParams{
		ID:    int64(req.GetId()),
		Count: req.GetCount(),
	}
	rec, err := server.Store.Queries.UpdateStorehouseItem(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.UpdateStorehouseItemResponse{
		Result: &api.StorehouseItem{
			Id:           int32(rec.ID),
			ComponentId:  rec.ComponentID,
			StorehouseId: rec.StorehouseID,
			Count:        rec.Count,
		},
	}, nil
}

func (server *GRPCServer) DeleteStorehouse(ctx context.Context, req *api.DeleteStorehouseRequest) (*api.DeleteStorehouseResponse, error) {
	err := server.Store.Queries.DeleteStorehouse(ctx, int64(req.GetId()))

	if err != nil {
		return nil, err
	}

	return &api.DeleteStorehouseResponse{}, nil
}

func (server *GRPCServer) DeleteStorehouseItem(ctx context.Context, req *api.DeleteStorehouseItemRequest) (*api.DeleteStorehouseItemResponse, error) {
	err := server.Store.Queries.DeleteItem(ctx, int64(req.GetId()))

	if err != nil {
		return nil, err
	}

	return &api.DeleteStorehouseItemResponse{}, nil
}

func (server *GRPCServer) DeleteStorehouseItemsByComponentIds(ctx context.Context, req *api.DeleteStorehouseItemsByComponentIdsRequest) (*api.DeleteStorehouseItemResponse, error) {
	params := sqlc.DeleteItemsByStorehouseAndComponentIdsParams{
		ComponentsIds: req.ComponentIds,
		StorehouseID:  req.StorehouseId,
	}
	err := server.Store.Queries.DeleteItemsByStorehouseAndComponentIds(ctx, params)

	if err != nil {
		return nil, err
	}

	return &api.DeleteStorehouseItemResponse{}, nil
}
