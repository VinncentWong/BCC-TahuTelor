package db

import (
	"module/entities/admin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	_db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/bcc-tahutelor?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	db = _db
	err = db.AutoMigrate(admin.Admin{})
	return db, nil
}

func GetDB() *gorm.DB {
	return db
}
