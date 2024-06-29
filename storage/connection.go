package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "EcoTrack"
	password = "hamidjon4424"
)

func Connect() (*sql.DB, error){
	connector := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode = disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", connector)
	return db, err
}
