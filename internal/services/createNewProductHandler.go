package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type CreateNewProductArgs struct {
	Name        string
	Size        string
	Value       uint
	WarehouseID uint
}

func (s *WarehouseManagementService) CreateNewProductHandler(_ *http.Request, args *CreateNewProductArgs, response *Response) error {
	log.Println("Starting to creating new product with warehouse_id=" + strconv.FormatUint(uint64(args.WarehouseID), 10) +
		", product_name=" + args.Name)

	code, err := s.Product.CreateNewProduct(args.Name, args.Size, args.Value, args.WarehouseID)
	if err != nil {
		return fmt.Errorf("creation product in Postgres has failed; error:  %e", err)
	}

	response.Message = "Done"

	log.Println("The product has been created successful: product_code=" + code.String())

	return nil
}
