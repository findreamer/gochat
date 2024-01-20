package main

import (
	"gochat/docs"
	"gochat/router"
	"gochat/utils"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swragger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
