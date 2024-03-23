package controllers

import (
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	db := database.GetDatabase()
	var order models.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	order.Price = 0
	order.Status = "0"
	err = db.Create(&order).Error
	if err != nil {
		utils.Respfailed("Захиалга үүсгэж чадсангүй !!! ", c, err.Error())
		return
	}

	utils.RespSuccess(order, "", c)
}

func GetAllOrders(c *gin.Context) {
	db := database.GetDatabase()
	var p []models.Order
	err := db.Find(&p).Error

	if err != nil {
		utils.Respfailed("Cannot find User with id: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}

// Assuming the JSON body structure is something like { "ids": [1, 2, 3] }
type RequestBody struct {
	IDs []int `json:"ids"`
}

func DeleteOrders(c *gin.Context) {
	db := database.GetDatabase()

	var requestBody RequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.Respfailed("Invalid request body", c, err.Error())
		return
	}

	err := db.Delete(&models.Order{}, "id IN (?)", requestBody.IDs).Error

	if err != nil {
		utils.Respfailed("Cannot delete orders", c, err.Error())
		return
	}

	var p []models.Order
	err = db.Find(&p).Error
	if err != nil {
		utils.Respfailed("Cannot find User with id: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}
