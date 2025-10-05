package middlewares

import (
	"errors"
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

	// Status code defaults to 500
    code := fiber.StatusInternalServerError

    // Retrieve the custom status code if it's a *fiber.Error
    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    }

	error_map:=fiber.Map{
		"message": err.Error(),
		"status": code,
	}


	if h.cfg.DEBUG {
		error_map["stack"]=string(debug.Stack())
	}
	
	return c.Status(fiber.StatusInternalServerError).JSON(error_map)
}


type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Stack   string `json:"stack,omitempty"`
}
