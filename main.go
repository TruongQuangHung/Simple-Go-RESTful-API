package main

import (
	"tekdraw-bff/pkg"
	"tekdraw-bff/service"
	"tekdraw-bff/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	err := pkg.ConnectDatabase()
	utils.CheckErr(err)

	r := gin.Default()

	// API Blogs
	router := r.Group("/api/v1/users")
	{
		router.GET("/", service.GetUser)
		router.POST("/", service.AddUser)
	}

	// API Diagram
	router = r.Group("/api/v1/diagrams")
	{
		router.GET("/", service.GetDiagram)
		router.POST("/", service.AddDiagram)
		router.PUT("/:id", service.UpdateDiagram)    // Add route for updating a diagram
		router.DELETE("/:id", service.DeleteDiagram) // Add route for deleting a diagram
	}

	// By default it serves on :8080
	r.Run(":8080")
}
