package warehouseService

import (
	"errors"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/models"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage"
)

type WarehouseService struct {
	storage storage.Storage
}

func NewWarehouseService(storage storage.Storage) *WarehouseService {
	return &WarehouseService{
		storage: storage,
	}
}

func (ws *WarehouseService) CreateNewWarehouse(name string, isAvailable bool) (uint, error) {
	if name == "" {
		return 0, errors.New("warehouse name is empty")
	}

	id, err := ws.storage.CreateNewWarehouse(name, isAvailable)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ws *WarehouseService) GetWarehouseByID(ID uint) (*models.Warehouse, error) {
	wh, err := ws.storage.GetWarehouse(ID)
	if err != nil {
		return nil, err
	}

	return wh, nil
}
