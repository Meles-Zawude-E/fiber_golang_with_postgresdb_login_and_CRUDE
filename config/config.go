package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	//load env file
	err := godotenv.Load(".env")
	if err !=nil {
		fmt.Println("Not load env file in correct way")
	}
	return os.Getenv(key)
}