package controllers

import (
	"fmt"
	"strconv"
	"time"

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
	fmt.Printf("%v", p)
	if p.Exfired == "0" {
		p.Exfired = "2023-12-31 23:59:00"
	} else if p.Exfired == "1" {
		p.Exfired = "2024-06-01 23:59:00"
	}
	
	err = db.Create(&p).Error
	if err != nil {
		utils.Respfailed("Thesis үүсгэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}

func GetOwnThesis(c *gin.Context) {
	db := database.GetDatabase()
	teacher_id := c.Param("teacher_id")
	newid, err := strconv.Atoi(teacher_id)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	code := c.Param("code")
	fmt.Printf("code %s",code)
	var thesis []models.Thesis
	var currentTime = time.Now();

	if(code == "null"){
		err = db.Where("teacher_id = ? AND to_date(exfired, 'YYYY-MM-DD') > ?", newid, currentTime).Find(&thesis).Error
	}
	if(code == "1"){
		err = db.Where("teacher_id = ? AND to_date(exfired, 'YYYY-MM-DD') < to_date('2024-06-02', 'YYYY-MM-DD') and to_date(exfired, 'YYYY-MM-DD') > to_date('2023-12-31', 'YYYY-MM-DD')", newid).Find(&thesis).Error
	}
	if(code == "0"){
		err = db.Where("teacher_id = ? AND to_date(exfired, 'YYYY-MM-DD') <= to_date('2023-12-31', 'YYYY-MM-DD')", newid).Find(&thesis).Error
	}

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
	Where th.status = 2 AND th.teacher_id = ?  
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

		var applyThesisList []models.ApplyThesis
		err := db.Where("thesis_id = ?", id).Find(&applyThesisList).Error
		if err != nil {
			fmt.Print("Failed to fetch users:", err)
		}
		var ids []int
		for _, appthesis := range applyThesisList {
			ids = append(ids, int(appthesis.StudentId))
		}
		fmt.Printf("%+v", ids)
		var users []models.User
		err = db.Where("id in ?", ids).Find(&users).Error
		if err != nil {
			fmt.Print("Failed to fetch users:", err)
		}

		requestedThesis := dto.AllRequestedThesis{
			ID:              uint(id),
			Status:          uint(status),
			TeacherID:       uint(teacherID),
			MglName:         mglName,
			EngName:         engName,
			Content:         content,
			Requirement:     requirement,
			AppliedStudents: users,
		}

		allRequestedThesis = append(allRequestedThesis, requestedThesis)
	}
	utils.RespSuccess(allRequestedThesis, "", c)
}
