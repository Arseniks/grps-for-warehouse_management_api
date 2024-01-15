package services

import (
	"errors"
	"log"
	"strconv"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/services/models"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage"
	"github.com/google/uuid"
)

type ProductService struct {
	storage storage.Storage
}

func NewProductService(storage storage.Storage) *ProductService {
	return &ProductService{
		storage: storage,
	}
}

func (p *ProductService) CreateNewProduct(name string, size string, value uint, warehouseID uint) (*uuid.UUID, error) {
	if name == "" || size == "" || value == 0 || warehouseID == 0 {
		return nil, errors.New("invalid input fields")
	}

	productCode, err := p.storage.CreateNewProduct(name, size, value, uuid.New(), warehouseID)
	if err != nil {
		return nil, err
	}

	return productCode, nil
}

func (p *ProductService) GetProductByCode(code uuid.UUID) (*models.Product, error) {
	product, err := p.storage.GetProductByCode(code)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductService) GetAvailableProductsCount(warehouseID uint, code uuid.UUID) (uint, error) {
	count, err := p.storage.GetAvailableProductsCount(warehouseID, code)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (p *ProductService) ReservationProducts(codes []uuid.UUID, warehouseID uint, values []uint) error {
	if len(codes) != len(values) {
		return errors.New("different len product codes and values")
	}
	ok := make([]error, 0, len(codes))

	for i, val := range codes {
		err := p.storage.ReservationProduct(val, warehouseID, values[i])
		if err != nil {
			log.Println(
				"Failed canceling products in the amount of",
				strconv.FormatUint(uint64(values[i]), 10),
				"with product_code="+val.String(),
				"; error:", err,
			)

			ok = append(ok, err)
		}
	}

	if len(ok) != 0 {
		return errors.New("something wrong in reservation product")
	}

	return nil
}

func (p *ProductService) CancelProductReservation(codes []uuid.UUID, warehouseID uint, values []uint) error {
	if len(codes) != len(values) {
		return errors.New("different len product codes and values")
	}

	ok := make([]error, 0, len(codes))

	for i, val := range codes {
		err := p.storage.CancelProductReservation(val, warehouseID, values[i])
		if err != nil {
			log.Println(
				"Failed canceling products in the amount of",
				strconv.FormatUint(uint64(values[i]), 10),
				"with product_code="+val.String(),
				"; error:", err,
			)

			ok = append(ok, err)
		}
	}

	if len(ok) != 0 {
		return errors.New("something wrong in canceling reservation product")
	}

	return nil
}
