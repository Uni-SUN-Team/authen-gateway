package routes

import (
	"unisun/api/authen-listening/src/controllers"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(g *gin.RouterGroup) {
	g.POST("/token", controllers.Validate)
	g.POST("/call-signin", controllers.CallSignin)
	g.POST("/call-revoke", controllers.CallRevoke)
	g.POST("/call-check-refreshtoken", controllers.CallCheckRefreshToken)
	g.GET("/token-version/:id", controllers.GetTokenVersionById)
}
