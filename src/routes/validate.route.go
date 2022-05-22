package routes

import (
	"unisun/api/authen-listening/src/controllers"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(g *gin.RouterGroup) {
	g.POST("/token", controllers.Validate)
	g.POST("/call-signin", controllers.CallSignin)
	g.POST("/call-revoke")
	g.POST("/call-refresh")
}
