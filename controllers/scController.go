package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository/location_repository"
	"github.com/gitHomesPhp/repaircat/repository/sc_repository"
	"net/http"
)

func AddSc(context *gin.Context) {
	if canContinue(context) {
		location := entity.NewLocation(getLocationParams(context))
		err, location := location_repository.Flush(location)
		sc := entity.NewSc(getScParams(context))
		sc.AddLocation(location)
		err = sc_repository.Flush(sc)
		fmt.Println(err)
		context.JSON(http.StatusOK, sc.ToMap())
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}

func getScParams(ctx *gin.Context) (name, description, phone, email, site string) {
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
