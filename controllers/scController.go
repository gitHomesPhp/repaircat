package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/repository"
	"github.com/gitHomesPhp/repaircat/types"
	"net/http"
)

func GetSC(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func AddSc(context *gin.Context) {
	sc := types.BuildSc(getScParams(context))
	repository.AddSc(sc)

	context.JSON(http.StatusOK, sc)
}

func getScParams(ctx *gin.Context) (name string, description string, phone string, email string, site string) {
	name = ctx.PostForm("name")
	description = ctx.PostForm("description")
	phone = ctx.PostForm("phone")
	email = ctx.PostForm("email")
	site = ctx.PostForm("site")
	return
}
