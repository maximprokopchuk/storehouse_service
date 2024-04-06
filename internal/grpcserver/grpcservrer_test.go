package grpcserver_test

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/maximprokopchuk/storehouse_service/internal/config"
	"github.com/maximprokopchuk/storehouse_service/internal/grpcserver"
	"github.com/maximprokopchuk/storehouse_service/internal/store"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	cfg := config.NewConfig()
	_, err := toml.DecodeFile("../../configs/config.test.toml", cfg)
	assert.Nil(t, err)
	st := store.New(cfg.Store)
	server := grpcserver.New(st)
	assert.NotNil(t, server)
	assert.Equal(t, st, server.Store, "should include store")
	assert.NotNil(t, server.CreateStorehouse)
	assert.NotNil(t, server.GetStorehousesListByCityId)
	assert.NotNil(t, server.GetStorehouseItemById)
	assert.NotNil(t, server.GetStorehouseItemsByStorehouseId)
	assert.NotNil(t, server.GetStorehouseItemsByStorehouseIdAndComponentsIds)
	assert.NotNil(t, server.CreateStorehouseItemForStorehouse)
	assert.NotNil(t, server.UpdateStorehouseItem)
	assert.NotNil(t, server.DeleteStorehouseItem)
	assert.NotNil(t, server.DeleteStorehouseItemsByComponentIds)
	assert.NotNil(t, server.DeleteStorehouse)
}
