package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/controllers"
	"github.com/gitHomesPhp/repaircat/middleware"
)

func main() {
	route := gin.Default()
	route.Use(middleware.CORSMiddleware())
	route.LoadHTMLGlob("html/index.html")
	///////////////////////////////////////////////////////
	route.GET("/sc-list", controllers.ScList)
	route.POST("/sc", controllers.AddSc)
	route.GET("/sc/:id", controllers.GetSc)

	route.GET("/cities", controllers.CityList)

	route.GET("/city", controllers.AllCities)
	route.GET("/city/:id/underground", controllers.GetUndergroundByCity)
	route.POST("/city", controllers.AddCity)

	route.POST("/sc/:id/review", controllers.AddReview)
	route.GET("/sc/:id/reviews-info", controllers.GetReviewsInfo)

	///////////////////////////////////////////////////////
	route.Run()
}
