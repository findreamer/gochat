package main

import (
	"gochat/router"
	"gochat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
