package gorm

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConn struct {
	Instance *gorm.DB
	Error    error
}

var (
	conn          DBConn
	dbConnOnce    sync.Once
	GetConnection = getConnection
)

func UseDefaultConnection() {
	GetConnection = getConnection
}

func getConnection() *DBConn {
	dbConnOnce.Do(func() {
		var err error
		cfg := &gorm.Config{
			PrepareStmt: true,
		}

		host := os.Getenv("PAYMENT_GATEWAY_POSTGRES_HOST")
		port := os.Getenv("PAYMENT_GATEWAY_POSTGRES_PORT")
		user := os.Getenv("PAYMENT_GATEWAY_POSTGRES_USER")
		pwd := os.Getenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD")
		dbName := os.Getenv("PAYMENT_GATEWAY_POSTGRES_DBNAME")

		dsn := "host=%v user=%v password=%v dbname=%v port=%v sslmode=disable"
		db, err := gorm.Open(postgres.Open(fmt.Sprintf(dsn, host, user, pwd, dbName, port)), cfg)
		conn.Instance, conn.Error = db, err
	})
	return &conn
}
