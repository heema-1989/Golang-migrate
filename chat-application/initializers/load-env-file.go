package initializers

import (
	"chat-application/utils"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	envErr := godotenv.Load("./.env")
	utils.CheckError(envErr, "Error loading .env file")
}
