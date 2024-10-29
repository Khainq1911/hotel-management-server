package security

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenToken(claims *jwt.MapClaims, ctx echo.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
