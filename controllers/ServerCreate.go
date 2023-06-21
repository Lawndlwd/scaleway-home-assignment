package controllers

import (
	"example/homeAssignment/initializers"
	"example/homeAssignment/models"
	"fmt"

	"github.com/gin-gonic/gin"
)


func ServerCreate(ctx *gin.Context) {

	server := models.Server{Name: "Jinzhu", Type: 1, Status: 1}

	result := initializers.DB.Create(&server)
	if result.Error != nil {
		ctx.Status(400)
		fmt.Print(result.Error)
		return
	}

		ctx.JSON(200, gin.H{
			"server": server,
		})
	}