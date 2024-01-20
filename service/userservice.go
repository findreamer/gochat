package service

import (
	"gochat/models"
	"gochat/utils"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 用户列表
// @Success 200 {string} json {"code", "message"}
// @Router /user/getUerList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	utils.DB.Find(&data)

	c.JSON(200, gin.H{
		"message": data,
	})
}
