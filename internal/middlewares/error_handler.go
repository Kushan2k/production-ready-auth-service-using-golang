package middlewares

import (
	"github/go_auth_api/internal/config"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)


type ErrorHandler struct {
	cfg *config.Config
}

func NewErrorHandler(cfg *config.Config) *ErrorHandler {
	return &ErrorHandler{
		cfg: cfg,
	}
}

func (h *ErrorHandler) HandleError(c *fiber.Ctx, err error) error {

	error_map:=fiber.Map{
		"message": err.Error(),
		"status":fiber.StatusInternalServerError,
		
	}

	if h.cfg.DEBUG {
		error_map["stack"]=debug.Stack()
	}
	
	return c.Status(fiber.StatusInternalServerError).JSON(error_map)
}
