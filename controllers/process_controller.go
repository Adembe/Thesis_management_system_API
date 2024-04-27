package controllers

import (
	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/models"
	"go-rest-api/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)
func InsertProcess(c *gin.Context) {
	db := database.GetDatabase()
	var p dto.Process
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error binding Json " + err.Error(),
		})
		return
	}

	err = db.Model(&models.Thesis{}).Where("id = ?", p.ThesisID).Update("status", p.Status).Error 
	if err != nil {
		utils.Respfailed("update query failed: " ,c,err.Error())
        return
	}


	err = db.Model(&models.ApplyThesis{}).Where("thesis_id = ? AND student_id = ? AND teacher_id = ?", p.ThesisID, p.StudentID, p.TeacherID).Update("status", p.Status).Error 
	if err != nil {
		utils.Respfailed("update query failed: " ,c,err.Error())
        return
	}

	var process models.Process

	process.TeacherId = p.TeacherID
	process.StudentId = p.StudentID
	process.ThesisId = p.ThesisID
	process.ProcessStatus = 0
	process.Process1 = 0
	process.Process2 = 0
	process.Process3 = 0
	process.Process4 = 0


	err = db.Save(&process).Error
	if err != nil {
		utils.Respfailed("Database query failed: " ,c,err.Error())
        return
	}
	utils.RespSuccess(nil, "", c)
}


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

func GetProcessStudent(c *gin.Context){
	db := database.GetDatabase()
	var p []dto.ProcessThesis
	studentID := c.Param("student_id")
    newID, err := strconv.Atoi(studentID)
    if err != nil {
        utils.Respfailed("Error converting student ID from string to integer: ", c, err.Error())
        return
    }
	query := `
		SELECT p.id,p.teacher_id,p.student_id, p.thesis_id,p.process1,p.process2,p.process3,p.process4,p.process_status ,ut.lname as teacher_name, us.lname as student_name, us.programm as student_programm,th.mgl_name as thesis_name FROM processes p
		left join users ut on p.teacher_id = ut.id
		left join users us on p.student_id = us.id
		left join theses th on p.thesis_id = th.id
		where p.student_id = ?
	`
	if err := db.Raw(query,newID).Scan(&p).Error; err != nil {
		log.Fatal("Query failed:", err)
	}
	utils.RespSuccess(p, "", c)
}