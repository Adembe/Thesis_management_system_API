package utils

import (
	"fmt"
	"go-rest-api/dto"

	"github.com/gin-gonic/gin"
)

func RespSuccess(body any, message string, c *gin.Context) {
	response := dto.Response{
		Message: func() string {
			if message != "" {
				return message
			}
			return "Амжилттай"
		}(),
		Status: true,
		Body:   body,
	}
	c.JSON(200, response)
}

func Respfailed(message string, c *gin.Context, err string) {
	fmt.Println(gin.H{
		"Error": message + err,
	}) // write to log server

	response := dto.Response{
		Message: func() string {
			if message != "" {
				return message
			}
			return "Амжилтгүй"
		}(),
		Status: false,
		Body:   "",
	}
	c.JSON(200, response)
}
