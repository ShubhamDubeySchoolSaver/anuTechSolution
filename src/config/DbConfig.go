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

    // Log the connection details (without sensitive information like password)
    log.Printf("MYSQL_USER: %s", dbUser)
    log.Printf("MYSQL_HOST: %s", dbHost)
    log.Printf("MYSQL_PORT: %s", dbPort)
    log.Printf("MYSQL_DATABASE: %s", dbName)

    // Build the Data Source Name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    // Attempt to connect to the database
    db, err := sql.Open(dbDriver, dsn)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
        return nil, err
    }

    // Ping the database to ensure connection is valid
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping the database: %v", err)
        return nil, err
    }

    log.Println("Successfully connected to the MySQL database")
    dbconnslave = db

    return dbconnslave, nil
}
