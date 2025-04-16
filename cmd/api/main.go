package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rmrachmanfauzan/elpempek/internal/db"
	"github.com/rmrachmanfauzan/elpempek/internal/handler"
	"github.com/rmrachmanfauzan/elpempek/internal/repository"
)

func main()  {
	db.InitDB()
	db.RunMigrations()
	
	e := echo.New()


	userRepo := repository.NewUserRepository(db.DB)
	userHandler := handler.NewUserHandler(userRepo)


	userRoutes := e.Group("/users")
	userRoutes.POST("", userHandler.CreateUser)
	userRoutes.GET("/:id", userHandler.FindUser)

	
	e.Logger.Fatal(e.Start(":8080"))
}