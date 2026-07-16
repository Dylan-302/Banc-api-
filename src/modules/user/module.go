package user

import (
	common "banc-api/src/common/types"
	handlers "banc-api/src/modules/user/handlers"

	"go.uber.org/fx"
)

func ConfigureModuleRoutes(
	handlerCreateUsers *handlers.CreateUserHandler,
	handlerDeleteUsers *handlers.DeleteUserHandler,
	handlerFindByID *handlers.FindByIdUserHandler,
	handlerGetAllUsers *handlers.GetAllUsersHandler,
	handlerUpdateUser *handlers.UpdateUserHandler,
) {
	_ = common.SliceHandlers{
		Prefix: "/users",
		Routes: []common.HandlerModule{
			{
				Handler:      handlerCreateUsers.Handle,
				Route:        "/",
				Method:       "POST",
				RequiresAuth: true,
			},
			{
				Handler:      handlerDeleteUsers.Handle,
				Route:        "/:id",
				Method:       "DELETE",
				RequiresAuth: true,
			},
			{
				Handler:      handlerFindByID.Handle,
				Route:        "/:id",
				Method:       "GET",
				RequiresAuth: true,
			},
			{
				Handler:      handlerGetAllUsers.Handle,
				Route:        "/",
				Method:       "GET",
				RequiresAuth: true,
			},
			{
				Handler:      handlerUpdateUser.Handle,
				Route:        "/:id",
				Method:       "PUT",
				RequiresAuth: true,
			},
		},
	}
}

func ModuleProviders() []fx.Option {
	return []fx.Option{
		fx.Provide(handlers.NewCreateUserHandler),
		fx.Provide(handlers.NewDeleteUserHandler),
		fx.Provide(handlers.NewFindByIdUserHandler),
		fx.Provide(handlers.NewGetAllUsersHandler),
		fx.Provide(handlers.NewUpdateUserHandler),
		fx.Invoke(ConfigureModuleRoutes),
	}
}
