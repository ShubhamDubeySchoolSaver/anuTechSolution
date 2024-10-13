package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var dbconnslave *sql.DB = nil

func ConnectMySqlDbSlaveSingleton() (*sql.DB, error) {
	dbDriver := DBDRIVERSlave    // Database Driver Name
	dbUser := DBUSERSlave        // Database Username
	dbPassword := DBPASWORDSlave // Database  Password
	dbUrl := DBURLSlave          // Database ip/host with port
	dbName := DBNAMESlave        // Database Name
	if dbconnslave == nil {
		d, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@"+dbUrl+"/"+dbName)
		if err != nil {
			// panic(err.Error())
			return nil, err
		}
		dbconnslave = d
	}
	return dbconnslave, nil
}
