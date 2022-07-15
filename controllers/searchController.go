package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/services/searchservice"
	"net/http"
)

func MainSearch(ctx *gin.Context) {
	query := ctx.Query("q")
	city := ctx.Query("city")
	underground := ctx.Query("underground")
	municipality := ctx.Query("municipality")
	page := ctx.Query("page")

	searchService := searchservice.CreateService(query, city, underground, municipality, page)

	scCards := searchService.GetScCards()

	ctx.JSON(http.StatusOK, scCards)
}
