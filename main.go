package main

import (
	"context"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/Marie20767/go-web-app-template/api/routes"
	"github.com/Marie20767/go-web-app-template/internal/store"
	"github.com/Marie20767/go-web-app-template/internal/utils/config"
)

func run() error {
	ctx := context.Background()

	cfg, err := config.ParseEnv()
	if err != nil {
		return err
	}

	db, err := store.NewStore(ctx, cfg)
	if err != nil {
		return err
	}
	defer db.Close() //nolint:errcheck
	log.Println("connected to DB successfully!")

	e := echo.New()
	routes.RegisterAll(e, db)
	return e.Start(":" + cfg.Port)
}

func main() {
	if err := run(); err != nil {
		log.Println("server closed: ", err)
		os.Exit(1)
	}
}
