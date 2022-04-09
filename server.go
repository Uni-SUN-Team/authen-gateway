package main

import (
	"os"
	"unisun/api/authen-listening/src"
	"unisun/api/authen-listening/src/config"
)

func main() {
	config.SetENV()
	r := src.App()
	port := os.Getenv("PORT")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
