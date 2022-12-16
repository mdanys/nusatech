package main

import (
	"fmt"
	"nusatech/config"
	"nusatech/features/users/delivery"
	"nusatech/features/users/repository"
	"nusatech/features/users/service"
	"nusatech/utils/databases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := databases.InitDB(cfg)
	databases.MigrateDB(db)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	uRepo := repository.New(db)
	uService := service.New(uRepo)
	delivery.New(e, uService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.ServerPort)))
}
