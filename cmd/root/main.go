package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/Arseniks/jsonrpc_warehouse_management_api/config"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/app"
	"github.com/Arseniks/jsonrpc_warehouse_management_api/internal/storage"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.LoadConfig()

	conn, err := storage.GetDBConnection(&storage.PostgresConfig{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
		SSLMode:  cfg.SSLMode,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	newApp := app.NewApp(ctx, conn)

	newApp.Run()
}
