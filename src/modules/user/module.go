package user

import (
	"banc-api/src/modules/user/application/usecase"
	"banc-api/src/modules/user/domain/repositories"
	"banc-api/src/modules/user/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes inicializa y registra todas las rutas del módulo de usuarios
func RegisterRoutes(router *gin.RouterGroup, repo repositories.UserRepository) {
	// Casos de Uso
	getAllUC := usecase.NewGetAllUserUseCase(repo)
	findByIDUC := usecase.NewFindByIDUseCase(repo)
	createUserUC := usecase.NewCreateUserUseCase(repo)
	updateUserUC := usecase.NewUpdateUserUseCase(repo)
	deleteUserUC := usecase.NewDeleteUserUseCase(repo)

	// Handlers
	getAllHandler := handlers.NewGetAllUserHandler(getAllUC)
	findByIDHandler := handlers.NewFindByIDHandler(findByIDUC)
	createUserHandler := handlers.NewCreateUserHandler(createUserUC)
	updateUserHandler := handlers.NewUpdateUserHandler(updateUserUC)
	deleteUserHandler := handlers.NewDeleteUserHandler(deleteUserUC)

	// Rutas de Usuario
	users := router.Group("/users")
	{
		users.GET("", getAllHandler.Handle)
		users.GET("/:id", findByIDHandler.Handle)
		users.POST("", createUserHandler.Handle)
		users.PUT("/:id", updateUserHandler.Handle)
		users.DELETE("/:id", deleteUserHandler.Handle)
	}
}
