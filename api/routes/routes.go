package routes

import (
	"github.com/labstack/echo/v4"

	handlers "github.com/Marie20767/go-web-app-template/api/handlers/userhandler"
	"github.com/Marie20767/go-web-app-template/internal/store"
)

func RegisterAll(e *echo.Echo, db *store.Store) {
	userHandler := &handlers.UserHandler{DB: db}

	e.GET("/hello/:name", userHandler.Hello)
	// other routes here...
}
