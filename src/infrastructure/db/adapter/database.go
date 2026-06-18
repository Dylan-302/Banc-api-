package adapter

import (
	"banc-api/src/common/config"
	"banc-api/src/infrastructure/db/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect inicializa la conexión con PostgreSQL usando GORM
func Connect() (*gorm.DB, error) {
	dsn := config.GetDatabaseURL()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Ejecutar auto-migración del modelo User
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("Advertencia: No se pudo ejecutar la auto-migración: %v", err)
	}

	DB = db
	return db, nil
}
