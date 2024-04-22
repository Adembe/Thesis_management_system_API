package controllers

import (
	"fmt"
	"strconv"

	"go-rest-api/database"
	"go-rest-api/dto"
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


func GetAllRequested(c *gin.Context) {
	db := database.GetDatabase()
	teacher_id := c.Param("teacher_id")
	newid, err := strconv.Atoi(teacher_id)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var thesisList []map[string]interface{}
	var allRequestedThesis []dto.AllRequestedThesis

	query := `
	select * from theses th 
	left join (
		select ath.thesis_id,
			usr.id as user_id
		from apply_theses ath
		left join users usr on ath.student_id = usr.id
	) a on th.id = a.thesis_id 
	 
	Where th.teacher_id = ?
	`
	if err := db.Raw(query, newid).Scan(&thesisList).Error; err != nil {
		fmt.Print("Query failed:", err)
	}

	
	for _, thesis := range thesisList {
		
		// Extract fields from the map using type assertions
		id, _ := thesis["id"].(int64)
		status, _ := thesis["status"].(int64)
		teacherID, _ := thesis["teacher_id"].(int64)
		mglName, _ := thesis["mgl_name"].(string)
		engName, _ := thesis["eng_name"].(string)
		content, _ := thesis["content"].(string)
		requirement, _ := thesis["requirement"].(string)
		userID, _ := thesis["user_id"].(int64)


		fmt.Println("id:",userID)
		var users []models.User
		if userID != 0 {
			err := db.Where("id = ?", userID).Find(&users).Error
			if err != nil {
				fmt.Print("Failed to fetch users: %v\n", err)
			}
		}

		requestedThesis := dto.AllRequestedThesis{
			ID:             uint(id),
			Status:         uint(status),
			TeacherID:      uint(teacherID),
			MglName:        mglName,
			EngName:        engName,
			Content:        content,
			Requirement:    requirement,
			AppliedStudents: users,
		}

		allRequestedThesis = append(allRequestedThesis, requestedThesis)
	}

	
	if err != nil {
		utils.Respfailed("thesis авчрах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(allRequestedThesis, "", c)
}
