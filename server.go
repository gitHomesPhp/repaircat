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
	route.GET("/sc-list/search", controllers.MainSearch)
	route.POST("/sc", controllers.AddSc)
	route.GET("/sc/:id", controllers.GetSc)
	route.GET("/sc/:id/review", controllers.GetScReviews)

	route.GET("/cities", controllers.CityList)

	route.GET("/city", controllers.AllCities)
	route.GET("/city/:id/underground", controllers.GetUndergroundByCity)
	route.GET("/city/:id/municipality", controllers.GetMunicipalityByCity)
	route.POST("/city", controllers.AddCity)

	///////////////////////////////////////////////////////
	route.Run()
}
