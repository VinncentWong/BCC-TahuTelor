package admin

import (
	"module/entities/dto/admin_dto"
	"module/main"

	"github.com/gin-gonic/gin"
)

// Handler harus ada context biar ntar bisa kembalikan response dalam bentuk XML/JSON
func SignUpHandler(c *gin.Context) {
	var bodySignup admin_dto.SignUp
	err := c.BindJSON(&bodySignup)
	if err != nil {
		panic(err.Error())
	}
	db := main.GetDB()
	result := db.Create(&bodySignup)
	if result.Error != nil {
		panic(result.Error.Error())
	}
}

func LoginHandler(c *gin.Context) {

}
