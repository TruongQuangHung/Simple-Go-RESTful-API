package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleError(
	c *gin.Context,
	err error,
) {
	c.JSON(500, gin.H{"error": err.Error()})
}
