package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/elpempek/internal/router"
	"github.com/rmrachmanfauzan/elpempek/internal/db"
)

func main()  {
	db.InitDB()
	db.RunMigrations()
	
	e := echo.New()

	// Register all app routes
	router.RegisterUserRoutes(e, db.DB)

	
	e.Logger.Fatal(e.Start(":8080"))
}