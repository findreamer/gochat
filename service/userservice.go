package service

import (
	"gochat/models"
	"gochat/utils"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	utils.DB.Find(&data)

	c.JSON(200, gin.H{
		"message": data,
	})
}
