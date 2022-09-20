package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type jwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func HandleAuthJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Invalid token",
				"status":  http.StatusBadRequest,
			})
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, errors.New("Signing method invalid")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
		}

		userId := fmt.Sprintf("%v", claims["user_id"])

		c.Request().Header.Set("X-Header-UserId", userId)
		return next(c)
	}
}
