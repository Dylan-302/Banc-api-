package db

import (
	"fmt"
	"log"

	config "banc-api/src/common/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnection struct {
	*gorm.DB
}

func NewDBConnection(cfg *config.Config) *DBConnection {
	if cfg == nil || cfg.DB == nil {
		log.Println("database configuration is not available, using in-memory fallback")
		return &DBConnection{}
	}

	if cfg.DB.Host == "" || cfg.DB.Username == "" || cfg.DB.Dbname == "" || cfg.DB.Port == "" {
		log.Println("database configuration is incomplete, using in-memory fallback")
		return &DBConnection{}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn), // Solo warnings y errores
	})
	if err != nil {
		log.Printf("error connecting to database, using in-memory fallback: %s", err)
		return &DBConnection{}
	}

	return &DBConnection{db}
}
