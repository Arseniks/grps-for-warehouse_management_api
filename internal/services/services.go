package services

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/services/models"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage"
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
	productsTable := storage.NewPostgresDB(conn, "products")
	warehousesTable := storage.NewPostgresDB(conn, "warehouses")

	return &WarehouseManagementService{
		Product:   NewProductService(productsTable),
		Warehouse: NewWarehouseService(warehousesTable),
	}
}

type ProductArgs struct {
	Name        string
	Size        string
	Value       uint
	WarehouseID uint
	Code        uuid.UUID
}

type CreateNewProductArgs struct {
	Name        string
	Size        string
	Value       uint
	WarehouseID uint
}

type ProductResponse struct {
	Name        string
	Size        string
	Value       uint
	WarehouseID uint
	Code        uuid.UUID
}

type WarehouseArgs struct {
	Name        string
	IsAvailable bool
}

type GetWarehouseArgs struct {
	WarehouseID uint
}

type ReservationArgs struct {
	Codes       []uuid.UUID
	Value       []uint
	WarehouseID uint
}

type WarehouseResponse struct {
	ID          uint
	Name        string
	IsAvailable bool
}

type AvailableProductsCountResponse struct {
	Code        uuid.UUID
	Count       uint
	WarehouseID uint
}

type Response struct {
	Message string
}

func (s *WarehouseManagementService) CreateNewProduct(_ *http.Request, args *CreateNewProductArgs, response *Response) error {
	log.Println("Starting to creating new product with warehouse_id=" + strconv.FormatUint(uint64(args.WarehouseID), 10) +
		", product_name=" + args.Name)

	code, err := s.Product.CreateNewProduct(args.Name, args.Size, args.Value, args.WarehouseID)
	if err != nil {
		return fmt.Errorf("creation product in DB has failed; error:  %e", err)
	}

	response.Message = "Done"

	log.Println("The product has been created successful: product_code=" + code.String())

	return nil
}

func (s *WarehouseManagementService) GetProductByCode(_ *http.Request, args *ProductArgs, response *ProductResponse) error {
	log.Println("Starting to getting product by product_code=" + args.Code.String())

	product, err := s.Product.GetProductByCode(args.Code)
	if err != nil {
		return fmt.Errorf("error when getting good by id: %s", err)
	}

	response.Name = product.Name
	response.Code = product.Code
	response.Value = product.Value
	response.Size = product.Size
	response.WarehouseID = product.WarehouseID

	log.Println("Getting product by code has finished successful")

	return nil
}

func (s *WarehouseManagementService) ReservationProduct(_ *http.Request, args *ReservationArgs, response *Response) error {
	log.Println("Starting to reservation products")

	if err := s.Product.ReservationProducts(args.Codes, args.WarehouseID, args.Value); err != nil {
		return fmt.Errorf("reservation products has failed; error:  %e", err)
	}

	response.Message = "Done"

	log.Println("Reservation products has finished successful")

	return nil
}

func (s *WarehouseManagementService) CancelProductReservation(_ *http.Request, args *ReservationArgs, response *Response) error {
	log.Println("Starting to canceling products reservation")

	err := s.Product.CancelProductReservation(args.Codes, args.WarehouseID, args.Value)
	if err != nil {
		return fmt.Errorf("error when cancel reservation: %s", err)
	}

	response.Message = "Done"

	log.Println("Canceling products reservation has done successfully")

	return nil
}

func (s *WarehouseManagementService) GetAvailableProductsCount(_ *http.Request, args *ProductArgs, response *AvailableProductsCountResponse) error {
	log.Println("Starting to get unreserved products count with product_code=" + args.Code.String() + ", warehouse_id=" +
		strconv.FormatUint(uint64(args.WarehouseID), 10))

	count, err := s.Product.GetAvailableProductsCount(args.WarehouseID, args.Code)
	if err != nil {
		return fmt.Errorf("error when getting goods count by stock id: %s", err)
	}

	response.Code = args.Code
	response.WarehouseID = args.WarehouseID
	response.Count = count

	log.Println("Get unreserved products has finished successful")

	return nil
}

func (s *WarehouseManagementService) CreateNewWarehouse(_ *http.Request, args *WarehouseArgs, response *Response) error {
	log.Println("Starting create new warehouse " + args.Name)

	id, err := s.Warehouse.CreateNewWarehouse(args.Name, args.IsAvailable)
	if err != nil {
		err = fmt.Errorf("creation warehouse in DB has failed; error:  %e", err)

		log.Println(err)

		return err
	}

	response.Message = "Done"

	log.Println("The warehouse has been created successful with warehouse_id=" + strconv.FormatUint(uint64(id), 10))

	return nil
}

func (s *WarehouseManagementService) GetWarehouseByID(_ *http.Request, args *GetWarehouseArgs, response *WarehouseResponse) error {
	log.Println("Starting getting warehouse with warehouse_id=" + strconv.FormatUint(uint64(args.WarehouseID), 10))

	warehouse, err := s.Warehouse.GetWarehouseByID(args.WarehouseID)
	if err != nil {
		err = fmt.Errorf("getting warehouse from DB has failed; error:  %e", err)

		log.Println(err)

		return err
	}

	response.ID = warehouse.ID
	response.Name = warehouse.Name
	response.IsAvailable = warehouse.IsAvailable

	log.Println("Getting warehouse by ID has done successful")

	return nil
}
