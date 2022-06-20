package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func canContinue(ctx *gin.Context) bool {
	if ctx.PostForm("user") == "hui" && ctx.PostForm("token") == "bui" {
		fmt.Println("DAAAAAA")
		return true
	}

	fmt.Println("NOOOOOOOOOOO")

	return false
}
