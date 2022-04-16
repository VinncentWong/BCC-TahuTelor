package main

import (
	"module/controller"
	"module/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var anotherDB *gorm.DB

func main() {
	_db, err := db.InitDB()
	r := gin.Default()
	if err != nil {
		panic(err.Error())
	}
	anotherDB = _db
	controller.InitRouter(r)
	r.Run()
}
