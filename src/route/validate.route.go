package route

import (
	"unisun/api/authen-listening/src/controller"

	"github.com/gin-gonic/gin"
)

func ValidateJWT(g *gin.RouterGroup) {
	g.POST("/token", controller.Validate)
}
