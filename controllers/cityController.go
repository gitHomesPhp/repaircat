package controllers

import (
	"github.com/gin-gonic/gin"
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
