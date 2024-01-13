package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

const driverName = "postgres"

func GetDBConnection(cfg *PostgresConfig) (*sql.DB, error) {
	log.Printf("Starting connect to %s DB", driverName)

	db, err := sql.Open(
		driverName,
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host,
			cfg.Port,
			cfg.User,
			cfg.DBName,
			cfg.Password,
			cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("Connection established to db: %s", cfg.DBName))

	return db, nil
}
