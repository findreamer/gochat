package router

import (
	"gochat/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/index", service.GetIndex)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPassoword", service.FindUserByNameAndPassoword)

	// 发送消息
	r.GET("/user/sendMsg", service.SendMsg)

	return r
}
