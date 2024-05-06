package cmd

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConfigGet(envName string, defaultVal interface{}) string {
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}
	switch v := defaultVal.(type) {
	case func() string:
		return v()
	case string:
		return v
	default:
		return ""
	}
}

func DBConnect() (*sql.DB, error) {

	user := ConfigGet("POSTGRES_USER", "postgres")
	password := ConfigGet("POSTGRES_PASSWORD", nil)
	host := ConfigGet("HBK_DB_HOST", "localhost")
	dbname := ConfigGet("POSTGRES_DB", "postgres")

	// Assemble the connection string
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, password, host, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %v", err)
	}
	return db, nil
}

func DBQuery(queryFunc func(db *sql.DB) error) error {
	db, err := DBConnect()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	return queryFunc(db)
}
