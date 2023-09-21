package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitializeBIMSDatabase(dbname, username, password, dbhost, dbport string) (*sqlx.DB, error) {
	conn := username + ":" + password + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname
	bimsdb, err := sqlx.Connect("mysql", conn)

	if err != nil {
		return nil, fmt.Errorf("Error in initializing bims database: %s", err)
	}

	return bimsdb, nil
}
