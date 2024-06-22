package main

import (
	"Ramenpedia/initializers"
	"Ramenpedia/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate((&models.Member{}))
}
