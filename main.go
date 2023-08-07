package main

import (
	"log"

	"github.com/ikura-hamu/bot_ikura-hamu/src/conf"
	"github.com/ikura-hamu/bot_ikura-hamu/src/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	logger, err := conf.NewLogger()
	if err != nil {
		log.Printf("failed to init logger: %v", err)
		return
	}

	router.Setup(logger)
}
