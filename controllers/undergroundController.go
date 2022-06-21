package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/repository/underground_repository"
	"net/http"
	"strconv"
)

func GetCityUnderground(ctx *gin.Context) {
	cityId, _ := strconv.Atoi(ctx.Param("id"))

	undergrounds, err := underground_repository.GetUndergroundByCityId(cityId)

	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, undergrounds)
}
