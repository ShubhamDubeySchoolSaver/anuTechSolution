package config

import(
	"os"
)

var DBDRIVERSlave = "mysql"
var DBUSERSlave = os.Getenv("MYSQL_USER")
var DBPASWORDSlave = os.Getenv("MYSQL_PASSWORD")
var DBURLSlave = "tcp(host.docker.internal:3306)"
var DBNAMESlave = "skool_saver"


