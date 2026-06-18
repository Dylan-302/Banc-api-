package infraestructure

import (
	config "banc-api/src/common/config"
	types "banc-api/src/common/types"
	"context"
	"errors"
	"fmt"
	"net/http"

	middleware "banc-api/src/infraestructure/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
)

func setRoutesByModule(app *fiber.App, h *types.HandlersStore) {
	// Prefijo base para todas las rutas de la API
	api := app.Group("/api/v1")

	for _, handlerModule := range h.Handlers {
		route := api.Group("/" + handlerModule.Prefix)
		for _, routeItem := range handlerModule.Routes {
			// Configurar la ruta específica
			switch routeItem.Method {
			case http.MethodGet:
				if routeItem.RequiresAuth {
					route.Get(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Get(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPost:
				if routeItem.RequiresAuth {
					route.Post(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Post(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPut:
				if routeItem.RequiresAuth {
					route.Put(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Put(routeItem.Route, routeItem.Handler)
				}
			case http.MethodDelete:
				if routeItem.RequiresAuth {
					route.Delete(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Delete(routeItem.Route, routeItem.Handler)
				}
			case http.MethodPatch:
				if routeItem.RequiresAuth {
					route.Patch(routeItem.Route, middleware.JWTMiddleware, routeItem.Handler)
				} else {
					route.Patch(routeItem.Route, routeItem.Handler)
				}
			}
		}
	}
}
func errorHandler(c fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	c.Status(code).JSON(fiber.Map{
		"isError": true,
		"message": err.Error(),
	})
	return nil
}

func NewHttpFiberServer(lc fx.Lifecycle, h *types.HandlersStore, cfg *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})

	app.Use(cors.New())
	app.Use(logger.NewWithConfig())

	app.Get("/health", func(c fiber.Ctx) error {
		c.JSON(fiber.Map{"status": "ok"})
		return nil
	})

	setRoutesByModule(app, h)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("Starting fiber server on port " + cfg.App.Port)
			go app.Listen(":" + cfg.App.Port)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.Shutdown()
		},
	})

	return app
}
