package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbconnslave *sql.DB = nil

func ConnectMySqlDbSlaveSingleton() (*sql.DB, error) {
	// Use environment variables to fetch database connection details
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	// Log connection details (without the password)
	log.Printf("Attempting to connect to MySQL database on host: %s, port: %s, database: %s", dbHost, dbPort, dbName)

	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	if dbconnslave == nil {
		// Open the database connection
		db, err := sql.Open(dbDriver, dsn)
		if err != nil {
			// Log the error if the connection fails
			log.Printf("Failed to connect to the database: %v", err)
			return nil, err
		}

		// Test the connection
		err = db.Ping()
		if err != nil {
			// Log if the ping to the database fails
			log.Printf("Failed to ping the database: %v", err)
			return nil, err
		}

		// Log successful connection
		log.Println("Successfully connected to the MySQL database")
		dbconnslave = db
	}

	return dbconnslave, nil
}
