package controllers

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/models"
	"go-rest-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func GetProcessAll(c *gin.Context) {
    db := database.GetDatabase()
    if db == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection is nil"})
        return
    }
    var p []dto.ProcessThesis
    query := `
        SELECT p.id, p.teacher_id, p.student_id, p.thesis_id, p.process1, p.process2, p.process3, p.process4, p.process_status, ut.lname as teacher_name, us.lname as student_name, us.programm as student_programm, th.mgl_name as thesis_name FROM processes p
        left join users ut on p.teacher_id = ut.id
        left join users us on p.student_id = us.id
        left join theses th on p.thesis_id = th.id
    `
    if err := db.Raw(query).Scan(&p).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    utils.RespSuccess(p, "", c)
}


func GetProcessTeacher(c *gin.Context){
	db := database.GetDatabase()
	var p []dto.ProcessThesis
	teacherId := c.Param("teacher_id")
    newID, err := strconv.Atoi(teacherId)
    if err != nil {
        utils.Respfailed("Error converting student ID from string to integer: ", c, err.Error())
        return
    }
	query := `
		SELECT p.id,p.teacher_id,p.student_id, p.thesis_id,p.process1,p.process2,p.process3,p.process4,p.process_status ,ut.lname as teacher_name, us.lname as student_name, us.programm as student_programm,th.mgl_name as thesis_name FROM processes p
		left join users ut on p.teacher_id = ut.id
		left join users us on p.student_id = us.id
		left join theses th on p.thesis_id = th.id
		where p.teacher_id = ?
	`
	if err := db.Raw(query,newID).Scan(&p).Error; err != nil {
		log.Fatal("Query failed:", err)
	}
	utils.RespSuccess(p, "", c)
}


func UpdateProcessOne(c *gin.Context){
	db := database.GetDatabase()
	var p models.Process
	err := c.ShouldBindJSON(&p)
	fmt.Printf("ppppp %v",p)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	err = db.Model(&models.Process{}).Where("id = ?", p.ID).Update("process1", p.Process1).Error

	if err != nil {
		utils.Respfailed("thesis хадгалах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}


func UpdateProcessAll(c *gin.Context){
	db := database.GetDatabase()
	var p models.Process
	err := c.ShouldBindJSON(&p)
	fmt.Printf("ppppp %v",p)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	err = db.Model(&models.Process{}).Where("id = ?", p.ID).Updates(models.Process{Process2: p.Process2, Process3: p.Process3, Process4: p.Process4}).Error


	if err != nil {
		utils.Respfailed("thesis хадгалах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}