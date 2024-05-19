package controllers

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/models"
	"go-rest-api/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllthesis(c *gin.Context) {
	db := database.GetDatabase()
	var p []dto.AllThesis
	var err error

	code := c.Param("code")
	var query string

	baseQuery := `
	SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
		   us.fname, us.lname, us.email, us.phone_number, us.address
	FROM theses th
	LEFT JOIN users us ON th.teacher_id = us.id 
	`

	switch code {
	case "null":
		query = baseQuery
	case "1":
		query = baseQuery + "WHERE th.exfired BETWEEN '2024-01-01' AND '2024-06-02'"
	case "0":
		query = baseQuery + "WHERE th.exfired <= '2023-12-31'"
	default:
		utils.Respfailed("Invalid code provided", c, "Invalid code")
		return
	}

	err = db.Raw(query).Scan(&p).Error

	if err != nil {
		utils.Respfailed("thesis авчрах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(p, "", c)
}



func UpdateReqThesis(c *gin.Context) {
	db := database.GetDatabase()
	var p models.Thesis
	err := c.ShouldBindJSON(&p)

	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}
	fmt.Printf("before exfried %s\n",p.Exfired)
	startDate, _ := time.Parse("2006-01-02", "2024-02-01")
	p.Exfired = startDate.Format("2006-01-02");
	fmt.Printf("\n after exfried %s\n",p.Exfired)
	err = db.Save(&p).Error

	if err != nil {
		utils.Respfailed("thesis хадгалах үед алдаа гарлаа !!! ", c, err.Error())
		return
	}



	var rows []models.Thesis
	err = db.Find(&rows).Error
	if err != nil {
		utils.Respfailed("Алдаа гарлаа: !!! ", c, err.Error())
		return
	}
	utils.RespSuccess(rows, "", c)
}
