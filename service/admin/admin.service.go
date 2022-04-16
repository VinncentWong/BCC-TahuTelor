package admin

import (
	"module/db"
	"module/entities/dto/admin_dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler harus ada context biar ntar bisa kembalikan response dalam bentuk XML/JSON
func SignUpHandler(c *gin.Context) {
	var bodySignup admin_dto.SignUp
	err := c.BindJSON(&bodySignup)
	if err != nil {
		panic(err.Error())
	}
	_db := db.GetDB()
	result := _db.Create(&bodySignup)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"successfull": true,
		"message":     "success add your data into database !",
	})
}

func LoginHandler(c *gin.Context) {

}
