package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainSearch(ctx *gin.Context) {
	query := ctx.Query("q")

	ctx.JSON(http.StatusOK, query)
}
