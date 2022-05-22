package config

import (
	"log"
	"os"
	"strings"
	"unisun/api/authen-listening/src/entitys"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	str := []string{
		"host=" + os.Getenv("DB_HOST"),
		"user=" + os.Getenv("DB_USER"),
		"password=" + os.Getenv("DB_PASS"),
		"dbname=" + os.Getenv("DB_NAME"),
		"port=" + os.Getenv("DB_PORT"),
		"sslmode=" + os.Getenv("DB_SSL"),
		"TimeZone=" + os.Getenv("DB_TIMEZONE"),
	}
	dsn := strings.Join(str, " ")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}

	var UserAuthPermission entitys.UserAuthPermission
	database.AutoMigrate(&UserAuthPermission)
	DB = database
}
