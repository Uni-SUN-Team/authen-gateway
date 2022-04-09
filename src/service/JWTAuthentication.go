package service

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtServices struct {
	SecretKey string `json:"secretKey"`
	Issure    string `json:"issure"`
}

func JWTAuthService() JWTService {
	return &jwtServices{
		SecretKey: getSecretKey(),
		Issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte(service.SecretKey), nil
	})

}
