package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/Marie20767/go-web-app-template/api/handlers"
)

func RegisterAll(e *echo.Echo, h *handlers.Handler) {
	e.GET("/health", h.HealthCheck)
	e.GET("/item/:id", h.GetItem)
}
