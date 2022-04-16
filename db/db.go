package db

import (
	"module/entities/admin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/bcc-tahutelor?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(admin.Admin{})
	return db, nil
}
