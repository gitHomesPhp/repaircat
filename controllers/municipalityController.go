package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/domain/municipality/postgresql"
	"net/http"
	"strconv"
)

func GetMunicipalityByCity(ctx *gin.Context) {
	cityId, _ := strconv.Atoi(ctx.Param("id"))

	municipalityRepo := postgresql.GetRepo()

	_, municipalities := municipalityRepo.ListByCity(cityId)

	ctx.JSON(http.StatusOK, municipalities)
}
