package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/ddd/domain/scCard/postgresql"
	"net/http"
	"strconv"
)

func ScList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		fmt.Println(err)
	}

	scCardRepo := postgresql.GetRepo()
	_, list, next := scCardRepo.List(page)

	ctx.JSON(http.StatusOK, []any{
		list,
		next,
	})
}
