package controllers

import (
	"log"

	"go-rest-api/database"
	"go-rest-api/dto"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

func GetStudentThesis(c *gin.Context) {
	db := database.GetDatabase()
	var p []dto.AllThesis
	query := `
SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement,
       us.fname, us.lname, us.email, us.phone_number, us.address
FROM theses th
LEFT JOIN users us ON th.teacher_id = us.id 
WHERE th.status = 2
`
if err := db.Raw(query).Scan(&p).Error; err != nil {
	log.Fatal("Query failed:", err)
}
	// err := db.Raw("SELECT th.id, th.status, th.teacher_id, th.mgl_name, th.eng_name, th.content, th.requirement, us.fname, us.lname, us.email, us.phone_number, us.address FROM theses th left join users us on  th.id = us.id").Scan(&p).Error
	// fmt.Printf("%v", p)
	// if err != nil {
	// 	utils.Respfailed("бүх thesis авах үед алдаа гарлаа !!! ", c, err.Error())
	// 	return
	// }
	utils.RespSuccess(p, "", c)
}
