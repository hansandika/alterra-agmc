package util

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type jwtCustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userId uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwtCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return err
	}
	return err
}

func ValidateUser(userIdA int, userIdB int) error {
	if userIdA != userIdB {
		return errors.New("this action is unauthorized")
	}
	return nil
}
