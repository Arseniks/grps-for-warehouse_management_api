package services

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type GetWarehouseArgs struct {
	WarehouseID uint
}

type WarehouseResponse struct {
	ID          uint
	Name        string
	IsAvailable bool
}

func (s *WarehouseManagementService) GetWarehouseByIDHandler(_ *http.Request, args *GetWarehouseArgs, response *WarehouseResponse) error {
	log.Println("Starting getting warehouse with warehouse_id=" + strconv.FormatUint(uint64(args.WarehouseID), 10))

	warehouse, err := s.Warehouse.GetWarehouseByID(args.WarehouseID)
	if err != nil {
		err = fmt.Errorf("getting warehouse from Postgres has failed; error:  %e", err)

		log.Println(err)

		return err
	}

	response.ID = warehouse.ID
	response.Name = warehouse.Name
	response.IsAvailable = warehouse.IsAvailable

	log.Println("Getting warehouse by ID has done successful")

	return nil
}
