package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rmrachmanfauzan/elpempek/internal/handler"
	"github.com/rmrachmanfauzan/elpempek/internal/db"
)

func main()  {
	
	db.InitDB()
	r := gin.Default()
	r.GET("/users", handler.GetUsers)

	r.Run(":8080")
}