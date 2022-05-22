package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"unisun/api/authen-listening/src/entitys"
	"unisun/api/authen-listening/src/models"
	"unisun/api/authen-listening/src/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Validate(c *gin.Context) {
	body := &models.Validate{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic(err.Error())
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		log.Panic(err.Error())
	} else {
		err = nil
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		log.Panic(err.Error())
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

func CallSignin(c *gin.Context) {
	body := &models.Validate{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Panic(err.Error())
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		log.Panic(err.Error())
	} else {
		err = nil
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		log.Panic(err.Error())
	} else {
		err = nil
	}
	user_auth_permission := entitys.UserAuthPermission{}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		if id, err := strconv.Atoi(claims["id"].(string)); err == nil {
			user_auth_permission.UserId = id
			user_auth_permission.Token = body.Token
			if iat, err := strconv.Atoi(claims["iat"].(string)); err == nil {
				user_auth_permission.Iat = iat
			} else {
				log.Panic(err)
			}
		} else {
			log.Panic(err)
		}
	}

}

// func TestRedirect(c *gin.Context) {
// 	http.Redirect(c.Writer, c.Request, "http://www.google.com", 301)
// }
