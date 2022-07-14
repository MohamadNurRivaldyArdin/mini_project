package middleware

import (
	"mini_project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int, Nama string) (string, error) {
	claims := jwt.MapClaims{}
	claims["Id"] = id
	claims["Nama"] = Nama
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
