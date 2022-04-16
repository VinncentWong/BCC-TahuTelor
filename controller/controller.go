package controller

import (
	"module/service/admin"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.POST("api/admin/signup", admin.SignUpHandler)
	r.POST("api/admin/login", admin.LoginHandler)
}
