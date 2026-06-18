package server

import (
	"banc-api/src/infrastructure/db/adapter"
	"banc-api/src/infrastructure/db/dao"
	"banc-api/src/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Factory maneja la inicialización de dependencias
type Factory struct {
	db *gorm.DB
}

// NewFactory crea un Factory conectándose a la base de datos
func NewFactory() (*Factory, error) {
	db, err := adapter.Connect()
	if err != nil {
		return nil, err
	}
	return &Factory{db: db}, nil
}

// SetupRoutes registra las rutas de todos los módulos del sistema
func (f *Factory) SetupRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")

	// Repositorios
	userRepo := dao.NewUserDAO(f.db)

	// Módulos
	user.RegisterRoutes(api, userRepo)
}
