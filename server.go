package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/controllers"
)

func main() {
	r := gin.Default()
	r.GET("/sc-list", controllers.GetScList)
	r.GET("/ping", controllers.PongController)
	r.GET("/sc/:id", controllers.GetSC)
	r.Run()
}
