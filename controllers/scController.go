package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gitHomesPhp/repaircat/entity"
	"github.com/gitHomesPhp/repaircat/repository/location_repository"
	"github.com/gitHomesPhp/repaircat/repository/sc_repository"
	"net/http"
	"strconv"
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

func GetSc(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	sc, _ := sc_repository.Find(id)

	ctx.JSON(http.StatusOK, sc)
}

func getScParams(ctx *gin.Context) (name, description, phone, email, site string) {
	name = ctx.PostForm("name")
	description = ctx.PostForm("description")
	phone = ctx.PostForm("phone")
	email = ctx.PostForm("email")
	site = ctx.PostForm("site")
	return
}

func getLocationParams(ctx *gin.Context) (
	city *entity.City,
	address string,
	undergrounds []*entity.Underground,
	latitude string,
	longitude string,
) {
	cityId, _ := strconv.Atoi(ctx.PostForm("city"))
	city = entity.EmptyCity()
	city.SetId(cityId)

	address = ctx.PostForm("address")
	latitude = ctx.PostForm("latitude")
	longitude = ctx.PostForm("longitude")

	undergroundsIds := ctx.PostFormArray("undergrounds")

	for _, strUndergroundId := range undergroundsIds {
		underground := entity.EmptyUnderground()
		undergroundId, _ := strconv.Atoi(strUndergroundId)
		fmt.Println(undergroundId)
		underground.SetId(undergroundId)

		undergrounds = append(undergrounds, underground)
	}

	return
}
