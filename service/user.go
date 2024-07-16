package service

import (
	"tekdraw-bff/model"
	"tekdraw-bff/repo"
	"tekdraw-bff/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user, err := repo.GetUser(name, password)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "no user found"})
		}
		return
	}

	c.JSON(200, gin.H{"data": user})
}

func AddUser(c *gin.Context) {
	request := model.AddUserRequest{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	if request.Name == "" || request.Password == "" {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	id, err := repo.AddUser(request.Name, request.Password)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"msg": "success", "id": id})
}
