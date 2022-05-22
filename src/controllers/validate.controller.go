package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"unisun/api/authen-listening/src/entitys"
	"unisun/api/authen-listening/src/gorms"
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
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		log.Panic(err.Error())
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		err = nil
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		log.Panic(err.Error())
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		err = nil
	}
	user_auth_permission := entitys.UserAuthPermission{}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		for key, val := range claims {
			log.Println("val ==> ", val, "   key ==> ", key, "   ", reflect.TypeOf(val))
			switch key {
			case "id":
				if val != 0 {
					user_auth_permission.UserId = int(val.(float64))
				} else {
					c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Value is undefinde.", Result: map[string]string{"type": key, "confirm": "false"}})
					return
				}
			case "token_version":
				if val.(string) != "" {
					user_auth_permission.TokenVersion = val.(string)
				} else {
					c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Value is undefinde.", Result: map[string]string{"type": key, "confirm": "false"}})
					return
				}
			case "iat":
				if val != 0 {
					user_auth_permission.Iat = val.(float64)
				} else {
					c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Value is undefinde.", Result: map[string]string{"type": key, "confirm": "false"}})
					return
				}
			case "exp":
				if val != 0 {
					user_auth_permission.Ext = val.(float64)
				} else {
					c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Value is undefinde.", Result: map[string]string{"type": key, "confirm": "false"}})
					return
				}
			default:
				c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Not found value.", Result: map[string]string{"type": key, "confirm": "false"}})
				return
			}
		}
	}
	if !CheckUserTokenIsNull(user_auth_permission) {
		c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Not found all value.", Result: map[string]string{"confirm": "false"}})
		return
	}
	gorms.JWTAuthService().FindAndCreate(user_auth_permission)
	c.JSON(http.StatusOK, &models.Response{Result: map[string]string{"confirm": "true"}})
}

func CheckUserTokenIsNull(user_auth_permission entitys.UserAuthPermission) bool {
	return user_auth_permission.UserId != 0 && user_auth_permission.TokenVersion != "" && user_auth_permission.Ext != 0 && user_auth_permission.Iat != 0
}
