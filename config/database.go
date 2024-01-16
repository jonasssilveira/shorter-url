package config

import (
	"database/sql"
	"fmt"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:pass321@localhost:5431/encode_url?sslmode=disable"
)

func GetDBClient() *sql.DB {
	DB, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		_ = fmt.Errorf(err.Error())
		return nil
	}
	return DB
}
