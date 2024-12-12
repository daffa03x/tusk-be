package main

import (
	"net/http"
	"tusk/config"
	"tusk/models"

	"github.com/gin-gonic/gin"
)


func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	config.CreateOwnerAccount(db)

	// Router
	router := gin.Default()
	router.GET("/",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World")
	})

	router.Static("/attachments", "./attachments")
	router.Run("localhost:8080")
}