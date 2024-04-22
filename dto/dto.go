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
	AppliedStudents []models.User `json:"applied_studets" gorm:"column:applied_studets"`
}