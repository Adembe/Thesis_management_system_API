package controllers

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/models"
	"go-rest-api/utils"
	"io/ioutil"
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
	fmt.Print("student id : ",studentID)
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

func InsertProcessDetail(c *gin.Context){
	db := database.GetDatabase()

	if err := c.Request.ParseMultipartForm(10 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot parse form"})
		return
	}
	processIdStr := c.Request.FormValue("processId")
	processId, err := strconv.ParseUint(processIdStr , 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thesis ID"})
		return
	}

	// Read additional fields
	thesisIdStr  := c.Request.FormValue("thesisId")
	// Convert string to uint
	thesisId, err := strconv.ParseUint(thesisIdStr , 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thesis ID"})
		return
	}
	studentIdStr := c.Request.FormValue("studentId")
	studentId, err := strconv.ParseUint(studentIdStr , 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thesis ID"})
		return
	}
	fileNameStr := c.Request.FormValue("fileName")

	teacherIdStr := c.Request.FormValue("teacherId")
	teacherId, err := strconv.ParseUint(teacherIdStr , 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid thesis ID"})
		return
	}
	file, _, err := c.Request.FormFile("pdf")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "upload file error"})
		return
	}
	defer file.Close()

	// Read file data into a byte slice
	pdfData, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file data"})
		return
	}


	var prDetail models.ProcessDetail
	prDetail.ProcessId = uint(processId)
	prDetail.ThesisID = uint(thesisId)
	prDetail.StudentID = uint(studentId)
	prDetail.TeacherID = uint(teacherId)
	prDetail.FileName = fileNameStr
	prDetail.Pdf_data = pdfData


	// Save to database
	if err := db.Create(&prDetail).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save process detail"})
		return
	}
	
	utils.RespSuccess(nil, "", c)
}

func GetProcessDetail(c *gin.Context){
	db := database.GetDatabase()
	process_id := c.Param("process_id")


	newid, err := strconv.Atoi(process_id)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	var processdetail []models.ProcessDetail
	
	err = db.Where("process_id = ?", newid).Find(&processdetail).Error
	
	if err != nil {
		utils.Respfailed("processdetail авчрах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(processdetail, "", c)
}