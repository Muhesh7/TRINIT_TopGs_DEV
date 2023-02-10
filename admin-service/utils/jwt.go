package utils

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/topgs/trinit/admin-service/config"
	"github.com/topgs/trinit/admin-service/schemas"
	"gorm.io/gorm"
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetCurrentUserFromToken(userToken string, db *gorm.DB) (schemas.User, error) {
	var user schemas.User
	token, err := jwt.ParseWithClaims(userToken, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(config.JwtSecret), nil
		})
	if err != nil {
		return user, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		claimsEmail := claims.Email
		log.Println("USER_EMAIL", claimsEmail)
		err = db.First(&user, "email = ?", claimsEmail).Error
		log.Println(user)
		if err != nil {
			return user, err
		}
	}
	return user, nil

}
