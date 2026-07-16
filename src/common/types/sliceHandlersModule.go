package common

import (
	"github.com/gofiber/fiber/v3"
)

type HandlerModule struct {
	Handler      func(fiber.Ctx) error
	Route        string
	Method       string
	RequiresAuth bool
}

type SliceHandlers struct {
	Prefix string
	Routes []HandlerModule
}
type GlobalHandlers []SliceHandlers

type HandlersStore struct {
	Handlers []SliceHandlers
}

func NewHandlersStore() *HandlersStore {
	return &HandlersStore{}
}
