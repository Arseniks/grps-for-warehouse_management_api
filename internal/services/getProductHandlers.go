package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type ProductArgs struct {
	Name        string
	Size        string
	Value       uint
	WarehouseID uint
	Code        uuid.UUID
}

type ProductResponse struct {
	Name        string
	Size        string
	Value       uint
	Code        uuid.UUID
	WarehouseID uint
}

type AvailableProductsCountResponse struct {
	Code        uuid.UUID
	Count       uint
	WarehouseID uint
}

func (s *WarehouseManagementService) GetProductByCodeHandler(_ *http.Request, args *ProductArgs, response *ProductResponse) error {
	log.Println("Starting to getting product by product_code=" + args.Code.String())

	product, err := s.Product.GetProductByCode(args.Code)
	if err != nil {
		return fmt.Errorf("error when getting good by id: %s", err)
	}

	response.Name = product.Name
	response.Size = product.Size
	response.Value = product.Value
	response.Code = product.Code
	response.WarehouseID = product.WarehouseID

	log.Println("Getting product by code has finished successful")

	return nil
}

func (s *WarehouseManagementService) GetAvailableProductsCountHandler(_ *http.Request, args *ProductArgs, response *AvailableProductsCountResponse) error {
	log.Println("Starting to get unreserved products count with product_code=" + args.Code.String() + ", warehouse_id=" +
		strconv.FormatUint(uint64(args.WarehouseID), 10))

	count, err := s.Product.GetAvailableProductsCount(args.WarehouseID, args.Code)
	if err != nil {
		return fmt.Errorf("error when getting goods count by stock id: %s", err)
	}

	response.Code = args.Code
	response.Count = count
	response.WarehouseID = args.WarehouseID

	log.Println("Get unreserved products has finished successful")

	return nil
}
