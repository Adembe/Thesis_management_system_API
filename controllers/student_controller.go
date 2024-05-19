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

func GetStudentThesis(c *gin.Context) {
	db := database.GetDatabase()
	var p []dto.AllThesis
	currentTime := time.Now().Format("2006-01-02") // Format current time as a date string
	fmt.Printf("current %s\n", currentTime)
	
	query := `
	SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
		us.fname, us.lname, us.email, us.phone_number, us.address
	FROM theses th
	LEFT JOIN users us ON th.teacher_id = us.id 
	WHERE th.status = 2 and th.exfired > '2024-01-01'
	`
	
	if err := db.Raw(query).Scan(&p).Error; err != nil {
		utils.Respfailed("Query failed:", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}




func StudentReqThesis(c *gin.Context) {
	db := database.GetDatabase()
	var p models.ApplyThesis
	err := c.ShouldBindJSON(&p)
	fmt.Printf("%v",&p)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error binding Json " + err.Error(),
		})
		return
	}

    var student []models.ApplyThesis
    err = db.Where("student_id = ? and thesis_id = ? and teacher_id = ?",p.StudentId,p.ThesisId,p.TeacherId).Find(&student).Error
    if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error saving the User: " + err.Error(),
		})
		return
	}
    if len(student) > 0 {
        utils.Respfailed("Хүсэлт илгээсэн байна: " ,c,"")
        return
    }

	err = db.Save(&p).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Error saving the User: " + err.Error(),
		})
		return
	}
	utils.RespSuccess(p,"",c)
}

func GetStudentShowReq(c *gin.Context) {
    db := database.GetDatabase()
    var p []dto.AllThesis
    studentID := c.Param("student_id")
    newID, err := strconv.Atoi(studentID)
    if err != nil {
        utils.Respfailed("Error converting student ID from string to integer: ", c, err.Error())
        return
    }

    query := `
        SELECT 
            ath.id,
            ath.status,
            aa.mgl_name,
            aa.eng_name,
            aa.content,
            aa.requirement,
            aa.fname,
            aa.lname,
            aa.email,
            aa.phone_number,
            aa.address
        FROM apply_theses ath
        LEFT JOIN (
            SELECT 
                th.id,
                th.status,
                th.teacher_id,
                th.mgl_name,
                th.eng_name,
                th.content,
                th.requirement,
                us.fname,
                us.lname,
                us.email,
                us.phone_number,
                us.address
            FROM theses th
            LEFT JOIN users us ON th.teacher_id = us.id 
            WHERE th.status = 2 or th.status = 3 
        ) aa ON ath.thesis_id = aa.id
        WHERE ath.student_id = ?
    `

    if err := db.Raw(query, newID).Scan(&p).Error; err != nil {
        utils.Respfailed("Database query failed: " ,c,err.Error())
        return
    }

    utils.RespSuccess(p, "", c)
}
