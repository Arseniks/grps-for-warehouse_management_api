package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/models"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage"
	"github.com/google/uuid"
)

type Postgres struct {
	client    *sql.DB
	tableName string
}

func NewPostgresDB(client *sql.DB, tableName string) storage.Storage {
	postgresDB := Postgres{
		client:    client,
		tableName: tableName,
	}

	return &postgresDB
}

func (p *Postgres) CreateNewProduct(name string, size string, value uint, code uuid.UUID, warehouseID uint) (*uuid.UUID, error) {
	transaction, err := p.client.Begin()
	if err != nil {
		return nil, err
	}

	createItemQuery := fmt.Sprintf(
		`INSERT INTO %s (product_name, product_size, product_value, product_code, warehouse_id, product_reserved_value) 
		VALUES ($1, $2, $3, $4, $5, 0) RETURNING product_code`,
		p.tableName,
	)
	row := transaction.QueryRow(createItemQuery, name, size, value, code, warehouseID)

	var productCode uuid.UUID
	err = row.Scan(&productCode)
	if err != nil {
		rollbackErr := transaction.Rollback()
		if rollbackErr != nil {
			log.Println("Error while rolling back the transaction")

			return nil, rollbackErr
		}

		return nil, err
	}

	if err = transaction.Commit(); err != nil {
		return nil, err
	}

	return &productCode, nil
}

func (p *Postgres) GetProductByCode(code uuid.UUID) (*models.Product, error) {
	query := fmt.Sprintf(
		`SELECT product_code, product_name, product_size, product_value FROM %s WHERE product_code = $1`,
		p.tableName,
	)
	row := p.client.QueryRow(query, code)

	var product models.Product
	err := row.Scan(&product.Code, &product.Name, &product.Size, &product.Value)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *Postgres) ReservationProduct(code uuid.UUID, warehouseID uint, value uint) error {
	count, err := p.GetAvailableProductsCount(warehouseID, code)
	if err != nil {
		return err
	}

	if count < value {
		return errors.New("there are not enough products in warehouse")
	}

	transaction, err := p.client.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		`UPDATE %s SET product_reserved_value = (product_reserved_value + $1) 
        WHERE product_code = $2 AND warehouse_id = $3`,
		p.tableName,
	)
	_, err = transaction.Exec(query, value, code, warehouseID)
	if err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) CancelProductReservation(code uuid.UUID, warehouseID uint, value uint) error {
	count, err := p.getReservedProductsCount(warehouseID, code)
	if err != nil {
		return err
	}

	if count < value {
		return errors.New("there are not enough reserved products in warehouse to canceling")
	}

	transaction, err := p.client.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(
		`UPDATE %s SET product_reserved_value = (product_reserved_value - $1)
        WHERE product_code = $2 AND warehouse_id = $3`,
		p.tableName,
	)
	_, err = transaction.Exec(query, value, code, warehouseID)
	if err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) getReservedProductsCount(warehouseID uint, code uuid.UUID) (uint, error) {
	query := fmt.Sprintf(
		`SELECT product_reserved_value FROM %s
        INNER JOIN warehouses ON products.warehouse_id = warehouses.warehouse_id
		WHERE products.warehouse_id = $1 and warehouse_available and product_code = $2`,
		p.tableName,
	)

	var count uint
	err := p.client.QueryRow(query, warehouseID, code).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *Postgres) GetAvailableProductsCount(warehouseID uint, code uuid.UUID) (uint, error) {
	query := fmt.Sprintf(
		`SELECT (product_value - product_reserved_value) FROM %s
        INNER JOIN warehouses ON products.warehouse_id = warehouses.warehouse_id
		WHERE products.warehouse_id = $1 and warehouse_available and product_code = $2`,
		p.tableName,
	)

	var count uint
	err := p.client.QueryRow(query, warehouseID, code).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *Postgres) CreateNewWarehouse(name string, isAvailable bool) (uint, error) {
	transaction, err := p.client.Begin()
	if err != nil {
		return 0, err
	}

	createItemQuery := fmt.Sprintf(
		"INSERT INTO %s (warehouse_name, warehouse_available) VALUES ($1, $2) RETURNING warehouse_id",
		p.tableName,
	)
	row := transaction.QueryRow(createItemQuery, name, isAvailable)

	var warehouseID uint
	err = row.Scan(&warehouseID)
	if err != nil {
		rollbackErr := transaction.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}

		return 0, err
	}

	if err = transaction.Commit(); err != nil {
		return 0, err
	}

	return warehouseID, nil
}

func (p *Postgres) GetWarehouse(ID uint) (*models.Warehouse, error) {
	query := fmt.Sprintf(
		`SELECT warehouse_id, warehouse_name, warehouse_available FROM %s WHERE warehouse_id = $1`,
		p.tableName,
	)
	row := p.client.QueryRow(query, ID)

	var warehouse models.Warehouse
	err := row.Scan(
		&warehouse.ID,
		&warehouse.Name,
		&warehouse.IsAvailable,
	)
	if err != nil {
		return nil, err
	}

	return &warehouse, nil
}
