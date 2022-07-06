package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddReview(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": false,
	})
}
