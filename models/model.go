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
	Exfired     string 	   
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


type NotificationThesis struct{
	ID          uint           `json:"id" gorm:"primaryKey"` 
	FromUser	uint			`json:"from_user"`
	ToUser		uint			`json:"to_user"`
	Type 		string 			`json:"type"`
	ThesisId    uint			`json:"thesis_id"`
	Thesis      Thesis 			`gorm:"foreignKey:ThesisId"`
	CreatedAt   time.Time      
	UpdatedAt   time.Time      
}


type Process struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	TeacherId     uint		   `json:"teacher_id"`
	StudentId	  uint         `json:"student_id"`
	ThesisId      uint		   `json:"thesis_id"`
	Process1      uint         `json:"process1"`
	Process2      uint         `json:"process2"`
	Process3      uint         `json:"process3"`
	Process4      uint         `json:"process4"`
	ProcessStatus uint		   `json:"process_status"`
	CreatedAt     time.Time    `json:"created"`
	UpdatedAt     time.Time    `json:"updated"`
}



type ProcessDetail struct {
	ID            uint         `json:"id" gorm:"primaryKey"`
	ProcessId     uint		   `json:"process_id"`
	Process	  	  Process      `gorm:"foreignKey:ProcessId"`
	TeacherID     uint		   `json:"teacher_id"`
	StudentID     uint		   `json:"student_id"`
	ThesisID     uint		   `json:"thesis_id"`
	Feedback      string	   `json:"feedback"`
	Status        uint	   	   `json:"status"`
	Pdf_data      []byte    	`json:"pdf_detail"`
	FileName      string 		`json:"file_name"`
	CreatedAt     time.Time    `json:"created"`
	UpdatedAt     time.Time    `json:"updated"`
}
