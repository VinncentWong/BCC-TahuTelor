package middleware

import (
	"math/rand"
	"module/entities/dto/admin_dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func GenerateToken(adminLogin admin_dto.Login) (string, error) {
	var randomToken string = RandomString(20)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": adminLogin.Email,
		"exp":   time.Now().Add(12 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(randomToken))
	return tokenString, err
}

func ValidateToken() {

}
