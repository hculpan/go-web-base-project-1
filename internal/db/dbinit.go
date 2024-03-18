package db

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	// Update this whatever variables you need to retrieve to
	// initialize the db
	/*
		dbInit = os.Getenv("DB_INIT")
		if dbInit == "" {
			log.Fatal("missing DB_INIT in configuration file")
		}
	*/
}
