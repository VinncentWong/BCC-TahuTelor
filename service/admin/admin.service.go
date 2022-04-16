package admin

import (
	"module/db"
	"module/entities/admin"
	"module/entities/dto/admin_dto"
	"module/middleware"
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
	var bodyLogin admin_dto.Login

	err := c.BindJSON(&bodyLogin)
	if err != nil {
		panic(err.Error())
	}

	_db := db.GetDB()
	var adminLogin admin_dto.Login
	result := _db.Where("email = ?", bodyLogin.Email).Take(&adminLogin)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	var admin admin.Admin
	result = _db.Where("email = ?", admin.Email).Take(&admin)
	if result.Error != nil {
		panic(result.Error.Error())
	}
	if bodyLogin.Password == adminLogin.Password {
		token, err := middleware.GenerateToken(adminLogin)
		if err != nil {
			panic(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"successful": true,
			"message":    "successfully login !!!",
			"data": gin.H{
				"id":    admin.ID,
				"name":  admin.Username,
				"email": admin.Email,
				"token": token,
			},
		})
	}
}
