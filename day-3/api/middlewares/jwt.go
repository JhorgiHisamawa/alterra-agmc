package middlewares

import (
	"time"

	"api/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config("SECRET_JWT")))
}

func ExtractTokenUseRole(e echo.Context) string {
	user := e.Get("role").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		return userRole
	}
	return "token is invalid"
}
