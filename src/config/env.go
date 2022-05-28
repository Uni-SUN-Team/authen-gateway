package config

import (
	"os"
	"unisun/api/authen-listening/src/constants"
)

func SetENV() {
	os.Setenv(constants.JWT_SECRET, "aSiAZgPRmmw7gN7p9WeQxQ==")
	os.Setenv(constants.CONTEXT_PATH, "/authen-listening")
	os.Setenv(constants.PORT, "8081")
	/**
	* DB
	 */
	os.Setenv(constants.DB_HOST, "unisun.dynu.com")
	os.Setenv(constants.DB_NAME, "unisunauthdb")
	os.Setenv(constants.DB_USER, "urquhmotrdhwqg")
	os.Setenv(constants.DB_PASS, "efad4bb2169e67ddaa17c21aba5c76efc6a9daa6a06310949eba9a006bf258da")
	os.Setenv(constants.DB_PORT, "5432")
	os.Setenv(constants.DB_SSL, "disable")
	os.Setenv(constants.DB_TIMEZONE, "Asia/Bangkok")

	os.Setenv(constants.LOG_PATH, "/Users/ns/Documents/UniSUN/authen-gateway/tmp/app.log")
}
