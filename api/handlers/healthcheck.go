package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) HealthCheck(e echo.Context) error {
	dbStatus := "ok"
	err := h.Store.Ping(e.Request().Context())
	if err != nil {
		dbStatus = "unreachable"
	}

	return e.JSON(http.StatusOK, map[string]string{
		"status": "ok",
		"db":     dbStatus,
	})
}
