package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"unisun/api/authen-listening/src/model"
	"unisun/api/authen-listening/src/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Validate(c *gin.Context) {
	body := &model.Validate{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal("Read token is fail.")
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		log.Fatal("Convert json token is fail.")
	} else {
		err = nil
	}
	result, err := service.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		log.Fatal("Function validate JWT token is error.")
	} else {
		err = nil
	}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		c.JSON(http.StatusOK, gin.H{
			"status": ok,
			"claims": claims,
		})
	} else {
		c.JSON(http.StatusFound, gin.H{
			"status": ok,
			"claims": claims,
		})
	}
}
