package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"unisun/api/authen-listening/src/service"

	"github.com/dgrijalva/jwt-go"
)

type FormAuthenResponse struct {
	JWT string `json:"jwt"`
}

func TestValidateJWT() {
	postBody, _ := json.Marshal(map[string]string{
		"identifier": "narawich",
		"password":   "-Naras-CPE290821-",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:1337/api/auth/local", "application/json", responseBody)
	if err != nil {
		log.Fatal("Request is not pass.")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Values is not parse.")
	}
	jwtPayload := &FormAuthenResponse{}
	err = json.Unmarshal([]byte(body), jwtPayload)
	if err != nil {
		log.Fatal("Values is not parse.")
	}
	var token string = jwtPayload.JWT
	response, _ := service.JWTAuthService().ValidateToken(token)
	if claims, ok := response.Claims.(jwt.MapClaims); ok && response.Valid {
		fmt.Println(claims["exp"], ok)
	} else {
		fmt.Println(err)
	}
}
