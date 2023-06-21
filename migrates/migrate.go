package main

import (
	"example/homeAssignment/initializers"
	"example/homeAssignment/models"
)


func init(){
	initializers.LoadEnvFile()
	initializers.ConnectToDB()
}
func main(){
	initializers.DB.AutoMigrate(&models.Server{})
}