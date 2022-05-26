package main

import (
	"os"
	"unisun/api/authen-listening/src"
	"unisun/api/authen-listening/src/config"
	"unisun/api/authen-listening/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.SetENV()
	}
	config.ConnectDatabase()
	r := src.App()
	port := os.Getenv("PORT")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
