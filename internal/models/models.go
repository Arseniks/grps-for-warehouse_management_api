package models

import "github.com/google/uuid"

type Product struct {
	Name          string    `json:"product_name"`
	Size          string    `json:"product_size"`
	Code          uuid.UUID `json:"product_code"`
	Value         uint      `json:"product_value"`
	WarehouseID   uint      `json:"warehouse_id"`
	ReservedValue uint      `json:"product_reserved_value"`
}

type Warehouse struct {
	Name        string `json:"warehouse_name"`
	IsAvailable bool   `json:"warehouse_available"`
	ID          uint   `json:"warehouse_id"`
}
