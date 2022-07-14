package middlewares

import (
	"log"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ID int) string {
	info := jwt.MapClaims{}
	info["ID"] = ID
	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := auth.SignedString([]byte("R4h@s1A!"))
	if err != nil {
		log.Fatal("cannot generate key")
		return ""
	}

	return token
}
