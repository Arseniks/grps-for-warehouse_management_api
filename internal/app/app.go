package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

const port = "8000"

type App struct {
	ctx     context.Context
	storage *sql.DB
}

func NewApp(ctx context.Context, storage *sql.DB) *App {
	return &App{
		ctx:     ctx,
		storage: storage,
	}
}

func (a *App) Run() {
	log.Println("Starting JSON-RPC server")

	server := rpc.NewServer()

	server.RegisterCodec(json.NewCodec(), "application/json")
	err := server.RegisterService(service.NewWarehouseManagementService(), "")
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	router.Handle("/rpc", server)

	go func() {
		if err = http.ListenAndServe(":"+port, router); err != nil {
			log.Fatalf("Listen and serve: %v", err)
		}
	}()

	log.Printf("Listen and Serve on http://localhost:%s/rpc", port)

	select {
	case <-a.ctx.Done():
		log.Println("Terminating server: context canceled")
		return
	}
}
