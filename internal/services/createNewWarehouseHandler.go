package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type WarehouseArgs struct {
	Name        string
	IsAvailable bool
}

func (s *WarehouseManagementService) CreateNewWarehouseHandler(_ *http.Request, args *WarehouseArgs, response *Response) error {
	log.Println("Starting create new warehouse " + args.Name)

	id, err := s.Warehouse.CreateNewWarehouse(args.Name, args.IsAvailable)
	if err != nil {
		err = fmt.Errorf("creation warehouse in Postgres has failed; error:  %e", err)

		log.Println(err)

		return err
	}

	response.Message = "Done"

	log.Println("The warehouse has been created successful with warehouse_id=" + strconv.FormatUint(uint64(id), 10))

	return nil
}
