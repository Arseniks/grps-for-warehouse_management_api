package service

import "net/http"

type WarehouseManagementArgs struct {
	Who string
}

type WarehouseManagementReply struct {
	Message string
}

type WarehouseManagementService struct{}

func NewWarehouseManagementService() *WarehouseManagementService {
	return &WarehouseManagementService{}
}

func (h *WarehouseManagementService) Say(r *http.Request, args *WarehouseManagementArgs, reply *WarehouseManagementReply) error {
	reply.Message = "Hello, " + args.Who + "!"

	return nil
}
