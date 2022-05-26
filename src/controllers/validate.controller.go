package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"unisun/api/authen-listening/src/entitys"
	"unisun/api/authen-listening/src/gorms"
	"unisun/api/authen-listening/src/logging"
	"unisun/api/authen-listening/src/models"
	"unisun/api/authen-listening/src/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Validate(c *gin.Context) {
	body := &models.Validate{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		logging.Println("", err.Error())
	} else {
		err = nil
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		logging.Println("", err.Error())
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
	body := &models.Signin{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		err = nil
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		logging.Println("", err.Error())
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		err = nil
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		logging.Println("", err.Error())
		return
	} else {
		err = nil
	}
	user_auth_permission := entitys.UserAuthPermission{}
	if body.UserId != 0 {
		user_auth_permission.UserId = body.UserId
	}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		for key, val := range claims {
			switch key {
			case "token_version":
				if val != 0 {
					user_auth_permission.TokenVersion = int(val.(float64))
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
			}
		}
	}
	if !CheckUserTokenIsNull(user_auth_permission) {
		c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Not found all value.", Result: map[string]string{"confirm": "false"}})
		return
	}
	if user_auth_permission_find := gorms.JWTAuthService().FindbyUserid(user_auth_permission.UserId); user_auth_permission_find.UserId == 0 {
		gorms.JWTAuthService().Create(user_auth_permission)
	} else {
		gorms.JWTAuthService().UpdateVersionToken(user_auth_permission.TokenVersion, user_auth_permission)
	}
	c.JSON(http.StatusOK, &models.Response{Result: map[string]string{"confirm": "true"}})
}

func CheckUserTokenIsNull(user_auth_permission entitys.UserAuthPermission) bool {
	return user_auth_permission.TokenVersion != 0 && user_auth_permission.Ext != 0
}

func CallRevoke(c *gin.Context) {
	body := models.RevokeBody{}
	if payload, err := ioutil.ReadAll(c.Request.Body); err != nil {
		logging.Println("", err.Error())
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		return
	} else {
		if err = json.Unmarshal(payload, &body); err != nil {
			logging.Println("", err.Error())
			c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
			return
		}
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		c.JSON(http.StatusFound, &models.Response{Error: err.Error(), Result: map[string]string{"confirm": "false"}})
		logging.Println("", err.Error())
		return
	} else {
		err = nil
	}
	user_auth_permission := entitys.UserAuthPermission{}
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		for key, val := range claims {
			switch key {
			case "token_version":
				if val != 0 {
					user_auth_permission.TokenVersion = int(val.(float64)) + 1
				} else {
					c.JSON(http.StatusUnprocessableEntity, &models.Response{Error: "Value is undefinde.", Result: map[string]string{"type": key, "confirm": "false"}})
					return
				}
			case "uid":
				if val != 0 {
					user_auth_permission.UserId = int(val.(float64))
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
			}
		}
	}
	if user_auth_permission_find := gorms.JWTAuthService().FindbyUserid(user_auth_permission.UserId); user_auth_permission_find.UserId != 0 && user_auth_permission_find.TokenVersion == user_auth_permission.TokenVersion-1 {
		gorms.JWTAuthService().UpdateVersionToken(user_auth_permission.TokenVersion, user_auth_permission)
		c.JSON(http.StatusOK, &models.Response{Result: map[string]string{"confirm": "true"}})
	} else {
		c.JSON(http.StatusOK, &models.Response{Result: map[string]string{"confirm": "false"}})
	}

}

func GetTokenVersionById(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		logging.Println("", err.Error())
	}
	userPermisson := gorms.JWTAuthService().FindbyUserid(userId)
	c.JSON(http.StatusOK, userPermisson)
}

func CallCheckRefreshToken(c *gin.Context) {
	logging.Println("Start CallCheckRefreshToken", "")
	body := &models.Validate{}
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logging.Println("", err.Error())
	}
	err = json.Unmarshal([]byte(jsonData), body)
	if err != nil {
		logging.Println("", err.Error())
	}
	result, err := services.JWTAuthService().ValidateToken(body.Token)
	if err != nil {
		logging.Println("", err.Error())
	}
	versionToken, userId, iat, ext := services.JWTAuthService().MappingRefreshToken(result)
	JWT := &models.RefreshJWT{
		TokenVersion: versionToken,
		Uid:          userId,
		Iat:          iat,
		Ext:          ext,
	}
	if time.Now().Unix()/1000 > ext {
		logging.Println("Token is timeout", "")
		c.JSON(http.StatusOK, &models.ReponseRefreshToken{
			Status:  false,
			Message: "Token is timeout",
			Claims:  *JWT,
		})
	}
	userAuthPermission := gorms.JWTAuthService().FindbyUserid(userId)
	if versionToken != userAuthPermission.TokenVersion {
		logging.Println("Refresh token is invalid", "")
		c.JSON(http.StatusOK, &models.ReponseRefreshToken{
			Status:  false,
			Message: "Refresh token is invalid",
			Claims:  *JWT,
		})
		return
	}
	logging.Println("End CallCheckRefreshToken", "")
	c.JSON(http.StatusOK, &models.ReponseRefreshToken{
		Status:  true,
		Message: "Refresh token is valid",
		Claims:  *JWT,
	})
}
