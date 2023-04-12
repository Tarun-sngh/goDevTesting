package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() { //has to be capital letters, it works as export
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
