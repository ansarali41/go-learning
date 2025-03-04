package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectDB connects to MySQL database
func ConnectDB() (*sql.DB, error) {
	// MySQL connection details
	connectionString := "root:rootuser@tcp(localhost:3306)/todo_db"

	// Open connection to database
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
