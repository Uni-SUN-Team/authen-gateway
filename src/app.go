package src

import (
	"unisun/api/authen-listening/src/route"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	validate := api.Group("/validate")
	{
		route.ValidateJWT(validate)
	}
	return r
}
