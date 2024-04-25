package gorm

import (
	"log"
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
		dsn := "host=localhost user=usuario password=contraseña dbname=nombre_basedatos port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), cfg)
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
