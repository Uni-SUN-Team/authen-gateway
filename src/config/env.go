package config

import "os"

func SetENV() {
	os.Setenv("SECRET", "aSiAZgPRmmw7gN7p9WeQxQ")
	os.Setenv("CONTEXT_PATH", "/authen-listening")
}
