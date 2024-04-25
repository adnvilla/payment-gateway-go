package gorm

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbInstance    *gorm.DB
	dbConnOnce    sync.Once
	GetConnection = getConnection
)

func UseDefaultConnection() {
	GetConnection = getConnection
}

func getConnection() *gorm.DB {
	dbConnOnce.Do(func() {
		var err error
		cfg := &gorm.Config{
			PrepareStmt: true,
		}

		host := os.Getenv("PAYMENT_GATEWAY_POSTGRES_HOST")
		user := os.Getenv("PAYMENT_GATEWAY_POSTGRES_USER")
		pwd := os.Getenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD")
		dbName := os.Getenv("PAYMENT_GATEWAY_POSTGRES_DBNAME")
		env := os.Getenv("PAYMENT_GATEWAY_ENV")

		dsn := "host=%v user=%v password=%v dbname=%v_%v port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(fmt.Sprintf(dsn, host, user, pwd, dbName, env)), cfg)
		if err != nil {
			log.Fatal("error: gorm db not found:", err)
		}
		dbInstance = db
	})
	return dbInstance
}

func Close() error {
	return nil
}
