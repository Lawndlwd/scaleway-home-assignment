package controllers

import "github.com/gin-gonic/gin"


func ServerGet(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "go",
		})
	}