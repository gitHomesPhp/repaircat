package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/domain/underground/postgresql"
	"net/http"
	"strconv"
)

func GetUndergroundByCity(ctx *gin.Context) {
	cityId, _ := strconv.Atoi(ctx.Param("id"))

	undergroundRepo := postgresql.GetRepo()

	_, undergrounds := undergroundRepo.ListByCity(cityId)

	ctx.JSON(http.StatusOK, undergrounds)
}
