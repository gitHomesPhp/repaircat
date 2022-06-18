package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository"
	"github.com/gitHomesPhp/repaircat/repository/sc_repository"
	"github.com/gitHomesPhp/repaircat/types"
	"net/http"
)

func GetSC(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Test(c *gin.Context) {
	list, err := sc_repository.List(1)

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, list)
}

func AddSc2(context *gin.Context) {
	sc := entity.NewSc(getScParams(context))
	err := sc_repository.Flush(sc)
	fmt.Println(err)
	context.JSON(http.StatusOK, sc.ToMap())
}

func AddSc(context *gin.Context) {
	if canContinue(context) {
		location := types.BuildLocation(getLocationParams(context))
		repository.AddLocation(location)
		locationId := repository.GetLocationId(location)
		sc := types.BuildSc(getScParams(context))
		sc.SetLocationId(locationId)

		repository.AddSc(sc)

		context.JSON(http.StatusOK, sc)
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
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
	address = ctx.PostForm("address")
	underground = ctx.PostForm("underground")
	return
}

func canContinue(ctx *gin.Context) bool {
	if ctx.PostForm("user") == "hui" && ctx.PostForm("token") == "bui" {
		fmt.Println("DAAAAAA")
		return true
	}

	fmt.Println("NOOOOOOOOOOO")

	return false
}
