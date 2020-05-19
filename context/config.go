package context

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GoDotEnvVariable to get .env variables
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
