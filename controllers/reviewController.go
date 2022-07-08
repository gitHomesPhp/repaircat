package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/service"
	"net/http"
	"strconv"
)

func AddReview(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
	})
}

func GetReviewsInfo(ctx *gin.Context) {
	scId, _ := strconv.Atoi(ctx.Param("id"))

	ctx.JSON(http.StatusOK, service.GetScReviewsInfo(scId))
}
