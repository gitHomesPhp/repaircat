package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", controllers.PongController)
	r.Run()
}
