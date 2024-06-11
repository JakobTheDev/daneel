package main

import (
	"log"

	"github.com/JakobTheDev/daneel/cmd/cli"
	"github.com/JakobTheDev/daneel/internal/database"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	cli.Execute()
}
