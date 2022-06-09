package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/repository"
	"net/http"
	"strconv"
)

func GetScList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))

	scList := repository.GetScList(page)

	fmt.Println(scList[0].Description)

	c.JSON(http.StatusOK, scList)
}
