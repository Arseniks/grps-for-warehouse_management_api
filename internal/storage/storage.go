package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/services/models"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateNewProduct(name string, size string, value uint, code uuid.UUID, warehouseID uint) (*uuid.UUID, error)
	GetProductByCode(code uuid.UUID) (*models.Product, error)
	ReservationProduct(code uuid.UUID, warehouseID uint, value uint) error
	GetAvailableProductsCount(warehouseID uint, code uuid.UUID) (uint, error)
	CancelProductReservation(code uuid.UUID, warehouseID uint, value uint) error

	CreateNewWarehouse(name string, isAvailable bool) (uint, error)
	GetWarehouse(id uint) (*models.Warehouse, error)
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

const driverName = "postgres"

func GetDBConnection(cfg *PostgresConfig) (*sql.DB, error) {
	log.Printf("Starting connect to %s DB", driverName)

	db, err := sql.Open(
		driverName,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.DBName,
			cfg.Password,
			cfg.SSLMode,
		),
	)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("Connection established to db: %s", cfg.DBName))

	return db, nil
}
