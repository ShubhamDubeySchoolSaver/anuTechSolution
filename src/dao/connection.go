package dao

import "database/sql"

type DbConn struct {
	DB *sql.DB
}
