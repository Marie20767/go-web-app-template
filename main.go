package main

import (
	"log"
	"os"

	"github.com/Marie20767/go-web-app-template/api/routes"
	"github.com/Marie20767/go-web-app-template/internal/store"
	"github.com/Marie20767/go-web-app-template/internal/utils"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func run() error {
	c, err := utils.ParseEnv()
	if err != nil {
		return err
	}

	db, err := store.NewStore(c.DbURL)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Println("connected to DB successfully!")

	e := echo.New()
	routes.RegisterAll(e, db)
	return e.Start(":" + c.Port);
}

func main() {
	if err := run(); err != nil {
		log.Println("server closed: ", err)
		os.Exit(1)
	}
}
