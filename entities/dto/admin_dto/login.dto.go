package admin_dto

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUp struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	City        string `json:"string"`
	PhoneNumber string `json:"phonenumber"`
}
