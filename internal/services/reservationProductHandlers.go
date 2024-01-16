package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type ReservationArgs struct {
	Codes       []uuid.UUID
	Value       []uint
	WarehouseID uint
}

func (s *WarehouseManagementService) ReservationProductHandler(_ *http.Request, args *ReservationArgs, response *Response) error {
	log.Println("Starting to reservation products")

	if err := s.Product.ReservationProducts(args.Codes, args.WarehouseID, args.Value); err != nil {
		return fmt.Errorf("reservation products has failed; error:  %s", err)
	}

	response.Message = "Done"

	log.Println("Reservation products has finished successful")

	return nil
}

func (s *WarehouseManagementService) CancelProductReservationHandler(_ *http.Request, args *ReservationArgs, response *Response) error {
	log.Println("Starting to canceling products reservation")

	err := s.Product.CancelProductReservation(args.Codes, args.WarehouseID, args.Value)
	if err != nil {
		return fmt.Errorf("error when cancel reservation: %s", err)
	}

	response.Message = "Done"

	log.Println("Canceling products reservation has done successfully")

	return nil
}
