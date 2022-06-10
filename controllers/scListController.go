package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/repository"
	"net/http"
	"strconv"
)

func GetScList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	scList := repository.GetScList2(page)

	c.JSON(http.StatusOK, scList)
}
