package godeal

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	// TODO: encrypt and move into configuration file.
	dbType   = "mysql"
	user     = "godealuser"
	pass     = "GoDeal1234"
	protocol = "tcp"
	host     = "localhost"
	port     = "3306"
	database = "godeal"
	charset  = "utf8"
	source   = user + ":" + pass + "@" + protocol + "(" + host + ":" + port + ")/" +
		database + "?parseTime=true&charset=" + charset
)

var db *sql.DB = nil

func init() {
	var err error
	db, err = open()
	if err != nil {
		panic(err.Error())
	}
	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

}

func open() (*sql.DB, error) {
	return sql.Open(dbType, source)
}

func Db() *sql.DB {
	return db
}
