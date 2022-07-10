package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/domain/city/postgresql"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository/city_repository"
	"net/http"
)

func AddCity(ctx *gin.Context) {
	if canContinue(ctx) {
		code := ctx.PostForm("code")
		label := ctx.PostForm("label")
		city := entity.NewCity(label, code)
		city_repository.Flush(city)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}

func AllCities(ctx *gin.Context) {
	cities, err := city_repository.GetAllCities()
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, cities)
}

func CityList(ctx *gin.Context) {
	cityRepo := postgresql.GetRepo()

	err, cities := cityRepo.List()
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, cities)
}
