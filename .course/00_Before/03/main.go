package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	if _, err := time.Parse("2006-01-02", "9999-15-31"); err != nil {
		fmt.Println(err)
	}

	if _, err := strconv.ParseFloat("12a", 64); err != nil {
		fmt.Println(err)
	}

	connStr := " password=master host=localhost sslmode=disable user=postgres password=master dbname=boxfish"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer  rows.Close()
}

type aaa interface() {
	Error() string
}

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
kkkkkkkkkk
errors.New("cannot divide by zero")

fmt.Errorf("user does not exist: %s", "john")
