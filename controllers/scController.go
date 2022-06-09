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

func AddSc(c *gin.Context) {
	sc := types.NewSc()
	name := c.PostForm("name")
	description := c.PostForm("description")
	phone := c.PostForm("phone")
	sc.Name = name
	sc.Description = description
	sc.Phone = phone
	repository.AddSc(sc)
	c.String(200, "name is %s", sc.Name)
}
