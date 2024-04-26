package dto

import (
	"go-rest-api/models"
)

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
	Body    any    `json:"body"`
}

type AllThesis struct {
	ID          uint   `json:"id" gorm:"column:id"`
	Status      string `json:"status" gorm:"column:status"`
	TeacherID   uint   `json:"teacher_id" gorm:"column:teacher_id"`
	MglName     string `json:"mgl_name" gorm:"column:mgl_name"`
	EngName     string `json:"eng_name" gorm:"column:eng_name"`
	Content     string `json:"content" gorm:"column:content"`
	Requirement string `json:"requirement" gorm:"column:requirement"`
	FirstName   string `json:"fname" gorm:"column:fname"`
	LastName    string `json:"lname" gorm:"column:lname"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Address     string `json:"address" gorm:"column:address"`
}


type AllRequestedThesis struct {
	ID          uint   `json:"id" gorm:"column:id"`
	Status      uint `json:"status" gorm:"column:status"`
	TeacherID   uint   `json:"teacher_id" gorm:"column:teacher_id"`
	MglName     string `json:"mgl_name" gorm:"column:mgl_name"`
	EngName     string `json:"eng_name" gorm:"column:eng_name"`
	Content     string `json:"content" gorm:"column:content"`
	Requirement string `json:"requirement" gorm:"column:requirement"`
	FirstName   string `json:"fname" gorm:"column:fname"`
	LastName    string `json:"lname" gorm:"column:lname"`
	Email       string `json:"email" gorm:"column:email"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	Address     string `json:"address" gorm:"column:address"`
	AppliedStudents []models.User `json:"applied_studets" gorm:"column:applied_students"`
}


type Process struct {
	Status      uint    `json:"status" gorm:"column:status"`
	TeacherID   uint    `json:"teacher_id" gorm:"column:teacher_id"`
	StudentID   uint    `json:"student_id" gorm:"column:student_id"`
	ThesisID 	uint	`json:"thesis_id" gorm:"column:thesis_id"`
}


type ProcessThesis struct {
	ID          uint   `json:"id" gorm:"column:id"`
	Status      uint   `json:"process_status" gorm:"column:process_status"`
	TeacherID   uint   `json:"teacher_id" gorm:"column:teacher_id"`
	TeacherName   string   `json:"teacher_name" gorm:"column:teacher_name"`
	StudentID   uint   `json:"student_id" gorm:"column:student_id"`
	StudentName   string   `json:"student_name" gorm:"column:student_name"`
	ThesisID   uint   `json:"thesis_id" gorm:"column:thesis_id"`
	Thesisname   string   `json:"thesis_name" gorm:"column:thesis_name"`
	StudentProgramm   string `json:"student_programm" gorm:"column:student_programm"`
	Process1    uint `json:"process1" gorm:"column:process1"`
	Process2    uint `json:"process2" gorm:"column:process2"`
	Process3    uint `json:"process3" gorm:"column:process3"`
	Process4    uint `json:"process4" gorm:"column:process4"`
}




