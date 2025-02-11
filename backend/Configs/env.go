package Configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("🔗 URI de connexion à MongoDB :", os.Getenv("MONGO_URI"))
	return os.Getenv("MONGO_URI")
}
