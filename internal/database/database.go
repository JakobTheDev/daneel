package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func InitDB() error {
	var err error

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbServer := os.Getenv("DB_SERVER")
	dbDatabase := os.Getenv("DB_DATABASE")

	dbConnectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;TrustServerCertificate=True",
		dbServer, dbUser, dbPassword, 12345, dbDatabase)

	DB, err = sql.Open("mssql", dbConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	return DB.Ping()
}
