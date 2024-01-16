package services

import (
	"database/sql"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/models"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/services/productService"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/services/warehouseService"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage/postgres"
	"github.com/google/uuid"
)

type Product interface {
	CreateNewProduct(name string, size string, value uint, warehouseID uint) (*uuid.UUID, error)
	GetProductByCode(code uuid.UUID) (*models.Product, error)
	ReservationProducts(codes []uuid.UUID, warehouseID uint, values []uint) error
	CancelProductReservation(codes []uuid.UUID, warehouseID uint, values []uint) error
	GetAvailableProductsCount(warehouseID uint, code uuid.UUID) (uint, error)
}

type Warehouse interface {
	CreateNewWarehouse(name string, isAvailable bool) (uint, error)
	GetWarehouseByID(ID uint) (*models.Warehouse, error)
}

type WarehouseManagementService struct {
	Product
	Warehouse
}

func NewWarehouseManagementService(conn *sql.DB) *WarehouseManagementService {
	productsTable := postgres.NewPostgresDB(conn, "products")
	warehousesTable := postgres.NewPostgresDB(conn, "warehouses")

	return &WarehouseManagementService{
		Product:   productService.NewProductService(productsTable),
		Warehouse: warehouseService.NewWarehouseService(warehousesTable),
	}
}

type Response struct {
	Message string
}
