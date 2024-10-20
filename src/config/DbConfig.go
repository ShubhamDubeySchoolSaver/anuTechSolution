package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "os/exec"
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

    // Log the environment variables (without sensitive info like password)
    log.Printf("MYSQL_USER: %s", dbUser)
    log.Printf("MYSQL_HOST: %s", dbHost)
    log.Printf("MYSQL_PORT: %s", dbPort)
    log.Printf("MYSQL_DATABASE: %s", dbName)

    // Build the Data Source Name (DSN)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

    log.Printf("Attempting to connect to MySQL with DSN: %s", dsn)  // Log the DSN (again, no password)

    if dbconnslave == nil {
        db, err := sql.Open(dbDriver, dsn)
        if err != nil {
            log.Printf("Failed to connect to the database: %v", err)
            return nil, err
        }

        err = db.Ping()
        if err != nil {
            log.Printf("Failed to ping the database: %v", err)
            return nil, err
        }

        log.Println("Successfully connected to the MySQL database")
        dbconnslave = db

        // Run the SQL script after connection
        ExecuteSQLScript()
    }

    return dbconnslave, nil
}

func ExecuteSQLScript() {
    // Command to run the SQL script
    scriptPath := "/docker-entrypoint-initdb.d/data.sql"
    cmd := exec.Command("mysql", "-h", os.Getenv("MYSQL_HOST"), "-u", os.Getenv("MYSQL_USER"),
        fmt.Sprintf("-p%s", os.Getenv("MYSQL_PASSWORD")), os.Getenv("MYSQL_DATABASE"), "-e", fmt.Sprintf("source %s", scriptPath))

    // Run the command
    err := cmd.Run()
    if err != nil {
        log.Fatalf("Error executing SQL script: %v", err)
    } else {
        log.Println("SQL script executed successfully")
    }
}
