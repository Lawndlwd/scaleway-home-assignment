package main

import (
	"example/homeAssignment/controllers"
	"example/homeAssignment/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
	// initializers.DB.AutoMigrate(&models.Server{})
}
func main() {

	r := gin.Default()

	r.GET("/server", controllers.ServerGet )
	r.POST("/server", controllers.ServerCreate )
	r.GET("/servers", controllers.ServerList )

	r.Run()
}





