package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/services/searchservice"
	"net/http"
	"strconv"
)

func MainSearch(ctx *gin.Context) {
	query := ctx.Query("q")
	city := ctx.Query("city")
	underground := ctx.Query("underground")
	municipality := ctx.Query("municipality")
	page, _ := strconv.Atoi(ctx.Query("page"))

	searchService := searchservice.CreateService(query, city, underground, municipality, page)

	scCards, nextPrevious := searchService.GetScCards()

	ctx.JSON(http.StatusOK, []any{
		scCards,
		nextPrevious,
	})
}
