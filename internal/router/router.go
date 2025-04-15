package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/rmrachmanfauzan/elpempek/internal/service"
	repo "github.com/rmrachmanfauzan/elpempek/internal/repository"
	handlers"github.com/rmrachmanfauzan/elpempek/internal/handler"
)

func RegisterUserRoutes(e *echo.Echo, db *gorm.DB) {
	repo := repo.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	userRoutes := e.Group("/users")
	{
		userRoutes.POST("", handler.CreateUser)
		userRoutes.GET("/:id", handler.GetUserByID)
	}
}
