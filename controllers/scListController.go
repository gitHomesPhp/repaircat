package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/repository/sc_repository"
	"net/http"
	"strconv"
)

func ScList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		fmt.Println(err)
	}

	scList, err := sc_repository.List(page)
	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, scList)
}
