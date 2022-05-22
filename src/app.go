package src

import (
	"os"
	"unisun/api/authen-listening/src/routes"

	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	api := r.Group(os.Getenv("CONTEXT_PATH") + "/api")
	validate := api.Group("/validate")
	{
		routes.ValidateJWT(validate)
	}
	return r
}
