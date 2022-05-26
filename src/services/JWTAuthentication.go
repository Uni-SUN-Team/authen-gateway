package services

import (
	"fmt"
	"os"
	"unisun/api/authen-listening/src/constants"
	"unisun/api/authen-listening/src/logging"

	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	MappingRefreshToken(jwt *jwt.Token) (int, int, int64, int64)
	MappingToken(jwtBody *jwt.Token) (int, int64, int64)
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
	secret := os.Getenv(constants.JWT_SECRET)
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

func (service *jwtServices) MappingRefreshToken(jwtBody *jwt.Token) (int, int, int64, int64) {
	var versionToken, uid int = 0, 0
	var iat, ext int64 = 0, 0
	if claims, ok := jwtBody.Claims.(jwt.MapClaims); ok && jwtBody.Valid {
		for key, val := range claims {
			switch key {
			case "token_version":
				versionToken = int(val.(float64))
			case "uid":
				uid = int(val.(float64))
			case "iat":
				iat = int64(val.(float64))
			case "exp":
				ext = int64(val.(float64))
			}
		}
	} else {
		logging.Println("JWT Pase error.", "")
	}
	return versionToken, uid, iat, ext
}

func (service *jwtServices) MappingToken(jwtBody *jwt.Token) (int, int64, int64) {
	var id int = 0
	var iat, ext int64 = 0, 0
	if claims, ok := jwtBody.Claims.(jwt.MapClaims); ok && jwtBody.Valid {
		for key, val := range claims {
			switch key {
			case "id":
				id = int(val.(float64))
			case "iat":
				iat = int64(val.(float64))
			case "exp":
				ext = int64(val.(float64))
			}
		}
	} else {
		logging.Println("JWT Pase error.", "")
	}
	return id, iat, ext
}
