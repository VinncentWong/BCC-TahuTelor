package admin

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username    string `json:"username"`
	Phonenumber string `json:"phonenumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	City        string `json:"city"`
	Description string `json:"description"`
}
