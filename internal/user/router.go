package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", handler.CreateUser)
		userRoutes.GET("/:id", handler.GetUserByID)
	}
}
