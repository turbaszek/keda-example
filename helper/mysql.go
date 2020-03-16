package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

// Build connection str
func makeMySQLClient() (*sql.DB, error) {
	// Build connection str
	config := mysql.NewConfig()
	config.Addr = fmt.Sprintf("%s:%s", "mysql", "3306")
	config.User = "root"
	config.DBName = "mysql"
	config.Passwd = os.Getenv("MYSQL_PASSWORD")
	config.Net = "tcp"
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Insert queries to the MySQL database
func Insert() error {
	db, err := makeMySQLClient()
	if err != nil {
		return err
	}

	// Create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS test (state varchar(20))")
	if err != nil {
		return err
	}

	// Insert some values
	values := make([]string, 100)
	for i := range values {
		values[i] = "('QUEUED')"
	}
	query := fmt.Sprintf("INSERT INTO test VALUES %s", strings.Join(values, ","))
	_, err = db.Exec(query)
	return err
}

// Delete all records from test table
func Delete() error {
	db, err := makeMySQLClient()
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM test WHERE 1 = 1")
	return err
}
