package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSC(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}
