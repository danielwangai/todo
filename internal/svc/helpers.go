package svc

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type JWTClaim struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	jwt.StandardClaims
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(firstName, lastName, email string) (string, error) {
	jwtKey := []byte(os.Getenv("TODO_APP_JWT_SECRET_TOKEN"))
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
