package models

import (
	"time"
)

type User struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Type        uint      `json:"type"`
	PhoneNumber *string   `json:"phone_number"`
	Address     *string   `json:"address"`
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}

type Order struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TrackCode   string    `json:"track_code"`
	Status      string    `json:"status"`
	Price       int       `json:"price"`
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
	Description string    `json:"description"`
	UserId      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserId"`
}

type Thesis struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	ThesisType	uint		   `json:"thesis_type"`
	ThesisId	uint		   `json:"thesis_id"`
	TeacherId	uint		   `json:"teacher_id"`
	User 		User 		   `gorm:"foreignKey:TeacherId"`
	TeacherName string         `json:"teacher_name"`
	MglName     string         `json:"user_name"`
	EngName		string		   `json:"user_type"`
	Content     string         `json:"content"`
	Requirement string         `json:"requirement"`
	CreatedAt   time.Time      
	UpdatedAt   time.Time      
}


