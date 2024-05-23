package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
}

func NewConnection(config *Config) (*sql.DB, error) {
	// MySQL DSN format: "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.User, config.Password, config.Host, config.Port, config.DBName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
