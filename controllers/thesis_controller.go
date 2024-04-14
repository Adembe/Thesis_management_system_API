package controllers

import (
	"fmt"
	"strconv"

	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func GetThesis(c *gin.Context) {
	db := database.GetDatabase()
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var thesis models.Thesis
	err = db.First(&thesis, newid).Error
	if err != nil {
		utils.Respfailed("thesis үүсгэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(thesis, "", c)
}

func CreateThesis(c *gin.Context) {
	db := database.GetDatabase()
	
	var p models.Thesis

	err := c.ShouldBindJSON(&p)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	fmt.Printf("%v",p)
	err = db.Create(&p).Error
	if err != nil {
		utils.Respfailed("Thesis үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(nil, "", c)
}

func GetOwnThesis(c *gin.Context) {
	db := database.GetDatabase()
	teacher_id := c.Param("teacher_id")
	newid, err := strconv.Atoi(teacher_id)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var thesis []models.Thesis

	err = db.Where("teacher_id = ?", newid).Find(&thesis).Error
	
	if err != nil {
		utils.Respfailed("thesis авчрах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(thesis, "", c)
}


func DeleteThesis(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Thesis{}, newid).Error

	if err != nil {
		utils.Respfailed("thesis үстгах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var p []models.Thesis
	err = db.Find(&p).Error
	if err != nil {
		utils.Respfailed("Алдаа гарлаа: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}

func UpdateThesis(c *gin.Context) {
	db := database.GetDatabase()
	var p models.Thesis
	err := c.ShouldBindJSON(&p)

	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	err = db.Save(&p).Error

	if err != nil {
		utils.Respfailed("thesis хадгалах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var rows []models.Thesis
	err = db.Find(&rows).Error
	if err != nil {
		utils.Respfailed("Алдаа гарлаа: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(rows, "", c)
}
