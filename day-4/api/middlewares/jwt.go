package middlewares

import (
	"time"

	"api/config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(ID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config("SECRET_JWT")))
}

func ExtractTokenUserID(e echo.Context) float64 {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["id"]
		return userID.(float64)
	}
	return 0
}
