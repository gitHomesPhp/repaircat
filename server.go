package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/controllers"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/index.html")
	///////////////////////////////////////////////////////
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/sc-list", controllers.GetScList)
	r.GET("/sc/:id", controllers.GetSC)
	r.POST("/sc", controllers.AddSc)
	r.POST("/sc/:id", controllers.AddSc)
	r.Run()
}
