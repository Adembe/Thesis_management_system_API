package controllers

import (
	"strconv"

	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID has to be integer",
		})
		return
	}

	var User models.User
	err = db.First(&User, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "User not found: " + err.Error(),
		})

		return
	}
	c.JSON(200, User)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var p models.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error creating the User: " + err.Error(),
		})
		return
	}
	utils.RespSuccess(nil, "", c)
}

func GetAllUsers(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.User
	err := db.Find(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find User with id: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID Has to be an integer",
		})
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.User{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error deleting the User: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()
	var p models.User
	err := c.ShouldBindJSON(&p)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error binding Json " + err.Error(),
		})
		return
	}

	err = db.Save(&p).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error saving the User: " + err.Error(),
		})
		return
	}

	c.JSON(200, p)
}
