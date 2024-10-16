package config

import (
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var dbconnslave *sql.DB = nil

func ConnectMySqlDbSlaveSingleton() (*sql.DB, error) {
	// Get database connection details from environment variables
	dbDriver := "mysql"                                  // Database Driver Name (can be hardcoded)
	dbUser := os.Getenv("MYSQL_USER")                    // Database Username
	dbPassword := os.Getenv("MYSQL_PASSWORD")            // Database Password
	dbUrl := os.Getenv("MYSQL_HOST") + ":3306"           // Database Host (with port 3306)
	dbName := os.Getenv("MYSQL_DATABASE")                // Database Name
	
	// If the database connection is not established, create a new connection
	if dbconnslave == nil {
		d, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@tcp("+dbUrl+")/"+dbName)
		if err != nil {
			return nil, err
		}
		dbconnslave = d
	}
	return dbconnslave, nil
}
