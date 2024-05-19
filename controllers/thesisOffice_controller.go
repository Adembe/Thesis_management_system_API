package controllers

import (
	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func GetAllthesis(c *gin.Context) {
	db := database.GetDatabase()
	var p []dto.AllThesis
	var err error;
    code := c.Param("code")
	query := `
	SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
		   us.fname, us.lname, us.email, us.phone_number, us.address
	FROM theses th
	LEFT JOIN users us ON th.teacher_id = us.id 
	`
	queryA := `
	SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
		   us.fname, us.lname, us.email, us.phone_number, us.address
	FROM theses th
	LEFT JOIN users us ON th.teacher_id = us.id 
	Where to_timestamp(exfired, 'YYYY-MM-DD') <= to_timestamp('2023-12-31', 'YYYY-MM-DD')
	`
	queryS := `
	SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
		   us.fname, us.lname, us.email, us.phone_number, us.address
	FROM theses th
	LEFT JOIN users us ON th.teacher_id = us.id 
	Where to_timestamp(exfired, 'YYYY-MM-DD') < to_timestamp('2024-06-02', 'YYYY-MM-DD') and to_timestamp(exfired, 'YYYY-MM-DD') > to_timestamp('2023-12-31', 'YYYY-MM-DD')
	`
	if(code == "null"){
		err = db.Raw(query).Find(&p).Error
	}
	if(code == "1"){
		err = db.Raw(queryS).Find(&p).Error
	}
	if(code == "0"){
		err = db.Raw(queryA).Find(&p).Error
	}
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
