package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/domain/reviewForScCardExtension/postgresql"
	"net/http"
	"strconv"
)

func GetScReviews(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	repo := postgresql.GetRepo()

	_, reviews, _ := repo.ListBySc(page, id)

	ctx.JSON(http.StatusOK, reviews)
}
