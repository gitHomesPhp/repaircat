package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetScList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hui": "pui",
	})
}
