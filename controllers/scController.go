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
	location := types.BuildLocation(getLocationParams(context))
	repository.AddSc(sc)
	repository.AddLocation(location)
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

func getLocationParams(ctx *gin.Context) (city string, address string, underground string) {
	city = ctx.PostForm("city")
	address = ctx.PostForm("description")
	underground = ctx.PostForm("underground")
	return
}
