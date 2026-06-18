package server

import (
	"banc-api/src/common/config"
	"banc-api/src/infrastructure/http"
	"log"

	"github.com/gin-gonic/gin"
)

type App struct {
	engine *gin.Engine
}

// NewApp crea una nueva instancia de la aplicación
func NewApp() *App {
	// Cargar configuración
	config.Load()

	engine := gin.Default()

	// Aplicar Middleware global de CORS
	engine.Use(http.CORSMiddleware())

	// Inicializar dependencias y configurar rutas
	factory, err := NewFactory()
	if err != nil {
		log.Fatalf("Error al inicializar las dependencias del servidor: %v", err)
	}

	factory.SetupRoutes(engine)

	return &App{engine: engine}
}

// Run arranca el servidor HTTP en el puerto configurado
func (a *App) Run() error {
	port := config.GetServerPort()
	log.Printf("Servidor escuchando en el puerto: %s", port)
	return a.engine.Run(":" + port)
}
