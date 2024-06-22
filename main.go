package main

import (
	"Ramenpedia/controllers"
	"Ramenpedia/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.GET("/members/:id", controllers.GetMember)
	r.GET("/members", controllers.GetMembers)
	r.POST("/members", controllers.CreateMember)
	r.PATCH("/members/:id", controllers.UpdateMember)
	r.DELETE("/members/:id", controllers.DeleteMember)
	r.Run()
}
