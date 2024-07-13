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

	// By default it serves on :8080
	r.Run()
}
