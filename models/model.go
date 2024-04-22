package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Type        uint      `json:"type"`
	Programm    uint      `json:"programm"`
	PhoneNumber *string   `json:"phone_number"`
	Address     *string   `json:"address"`
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}

type Thesis struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Status		uint		   `json:"status"` 
	TeacherId   uint		   `json:"teacher_id"` 
	User   		User		   `gorm:"foreignKey:TeacherId"`
	MglName     string         `json:"mgl_name"`
	EngName		string		   `json:"eng_name"`
	Content     string         `json:"content"`
	Requirement string         `json:"requirement"`
	CreatedAt   time.Time      
	UpdatedAt   time.Time      
}


type ApplyThesis struct{
	ID          uint           `json:"id" gorm:"primaryKey"`
	Status		uint		   `json:"status"` 
	ThesisId    uint           `json:"thesis_id"`
	StudentId   uint		   `json:"student_id"` 
	TeacherId   uint		   `json:"teacher_id"` 
	CreatedAt   time.Time      
	UpdatedAt   time.Time      
}


type Process struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	TeacherId     string		 `json:"teacher_id"`
	TeacherName   string         `json:"teacher_name"`
	StudentId	  string         `json:"student_id"`
	StudentName   string         `json:"student_name"`
	BranchSchool  string		 `json:"branch_school"`
	Chair         string		 `json:"chair"`
	ThesisName    string		 `json:"thesis_name"`
	Process1      string         `json:"process1"`
	Process2      string         `json:"process2"`
	Process3      string         `json:"process3"`
	Process4      string         `json:"process4"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
}

