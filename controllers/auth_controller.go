package controllers

import (
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/utils"

	"github.com/gin-gonic/gin"
)

type AuthResp struct {
	UserId   uint
	Token    string
	Type     uint
	UserName string
}

func Login(c *gin.Context) {
	db := database.GetDatabase()
	var reqModel models.User
	var userModel models.User

	err := c.ShouldBindJSON(&reqModel)
	if err != nil {
		utils.Respfailed("Json хөрвүүлэх үед алдаа гарлаа !!! ", c, err.Error())
		return
	}

	result := db.Where("email = ? AND password = ?", reqModel.Email, reqModel.Password).Find(&userModel)

	if result.RowsAffected == 0 {
		utils.Respfailed("Нууц үг эсвэл мейл хаяг буруу байна !!! ", c, "no user")
		return
	}

	var resp AuthResp
	resp.Token = "my token"
	resp.UserId = userModel.ID
	resp.Type = userModel.Type
	resp.UserName = userModel.Fname
	// b, err := json.Marshal(resp)
	// if err != nil {
	//     fmt.Println(err)
	//     return
	// }
	utils.RespSuccess(resp, "", c)
}
