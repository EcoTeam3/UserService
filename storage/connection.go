package storage

import (
	"database/sql"
	"fmt"
	"userService/config"

	_ "github.com/lib/pq"
)


func Connect() (*sql.DB, error){
	config := config.Load()
	connector := fmt.Sprintf(`host = %s port = %d user = %s dbname = %s password = %s sslmode = disable`,
								config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_NAME, config.DB_PASSWORD)
	db, err := sql.Open("postgres", connector)
	return db, err
}
