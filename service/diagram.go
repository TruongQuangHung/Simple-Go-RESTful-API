package service

import (
	"strconv"

	"tekdraw-bff/model"
	"tekdraw-bff/repo"
	"tekdraw-bff/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDiagram(c *gin.Context) {
	id := c.Query("id")
	// parse id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}

	diagram, err := repo.GetDiagram(idInt)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "no diagram found"})
		}
		return
	}

	c.JSON(200, gin.H{"data": diagram})
}

func AddDiagram(c *gin.Context) {
	request := model.AddDiagramRequest{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	if request.UserId == 0 || request.Type == "" || request.Code == "" {
		c.JSON(400, gin.H{"error": "invalid params"})
		return
	}
	id, err := repo.AddDiagram(request.UserId, request.Type, request.Code)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"msg": "success", "id": id})
}

func UpdateDiagram(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid diagram ID"})
		return
	}

	var request model.UpdateDiagramRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	// Assuming repo.UpdateDiagram exists and updates the diagram based on the provided ID and request data
	if err := repo.UpdateDiagram(id, request); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"msg": "diagram updated"})
}

func DeleteDiagram(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid diagram ID"})
		return
	}

	// Assuming repo.DeleteDiagram exists and deletes the diagram based on the provided ID
	if err := repo.DeleteDiagram(id); err != nil {
		utils.HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{"msg": "diagram deleted"})
}
