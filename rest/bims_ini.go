package rest

import (
	"bims/database"
	"fmt"

	"github.com/jmoiron/sqlx"
	"gopkg.in/ini.v1"
)

type BimsConfiguration struct {
	BIMSdb     *sqlx.DB
	from       string
	pass       string
	FrontEndIP string
	PortNumber string
}

func New() (*BimsConfiguration, error) {

	// read config file
	cfg, err := ini.Load("/root/bims-backend/config.ini")
	if err != nil {
		return nil, fmt.Errorf("Fail to read file: %v", err)
	}

	dbSection := cfg.Section("db")
	user := dbSection.Key("user").String()
	password := dbSection.Key("password").String()
	dbhost := dbSection.Key("dbhost").String()
	dbport := dbSection.Key("dbport").String()
	dbname := dbSection.Key("dbname").String()

	emailSection := cfg.Section("email")
	from := emailSection.Key("from").String()
	pass := emailSection.Key("pass").String()

	Frontend := cfg.Section("frontend")
	ip := Frontend.Key("ip").String()
	port := Frontend.Key("port").String()

	bimsdb, err := database.InitializeBIMSDatabase(dbname, user, password, dbhost, dbport)
	if err != nil {
		return nil, err
	}

	return &BimsConfiguration{
		bimsdb,
		from,
		pass,
		ip,
		port,
	}, nil
}
