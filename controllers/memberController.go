package controllers

import (
	"Ramenpedia/initializers"
	"Ramenpedia/models"

	"github.com/gin-gonic/gin"
)

func GetMember(c *gin.Context) {
	id := c.Param("id")
	var member models.Member
	initializers.DB.First(&member, id)

	c.JSON(200, gin.H{
		"member": member,
	})
}

func GetMembers(c *gin.Context) {
	var members []models.Member
	initializers.DB.Find(&members)

	c.JSON(200, gin.H{
		"members": members,
	})
}

func CreateMember(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name     string
		Email    string
		Birthday string
		Token    string
	}

	c.Bind(&body)

	// Create a new member
	member := models.Member{
		Name:     body.Name,
		Email:    body.Email,
		Birthday: body.Birthday,
		Token:    body.Token,
	}

	result := initializers.DB.Create(&member)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"member": member,
	})
}

func UpdateMember(c *gin.Context) {
	id := c.Param("id")

	// Get data off req body
	var body struct {
		Name     string
		Email    string
		Birthday string
		Token    string
	}

	c.Bind(&body)

	var member models.Member
	initializers.DB.First(&member, id)

	initializers.DB.Model(&member).Updates(models.Member{
		Name:     body.Name,
		Email:    body.Email,
		Birthday: body.Birthday,
		Token:    body.Token,
	})

	c.JSON(200, gin.H{
		"member": member,
	})
}

func DeleteMember(c *gin.Context) {
	id := c.Param("id")
	var member models.Member
	initializers.DB.Delete(&member, id)

	c.JSON(200, gin.H{
		"member": member,
	})
}
