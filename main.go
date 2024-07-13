package main

import (
	"tekdraw-bff/controller"
	"tekdraw-bff/pkg"
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
		router.GET("/", controller.GetUser)
		router.POST("/", controller.AddUser)
	}

	// By default it serves on :8080
	r.Run()
}
