/*
Copyright © 2024 Jakob Pennington
*/
package main

import (
	"database/sql"
	"log"

	"github.com/JakobTheDev/daneel/cmd"
	"github.com/JakobTheDev/daneel/internal/database"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var db *sql.DB

type Platform struct {
	ID          int64
	DisplayName string
}

func main() {
	var err error

	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Execute()
}
