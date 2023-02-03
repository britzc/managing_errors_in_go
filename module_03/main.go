package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

func DateParseError() {
	if _, err := time.Parse("2006-01-02", "9999-15-31"); err != nil {
		fmt.Println(err)
	}
}

func NumberParseError() {
	if _, err := strconv.ParseFloat("12a", 64); err != nil {
		fmt.Println(err)
	}
}

func DatabaseError() {
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
	defer rows.Close()
}

func main() {
	heading("Date Parse Error Output")
	DateParseError()

	heading("Number Parse Error Output")
	NumberParseError()

	heading("Database Error Output")
	DatabaseError()
}

func heading(val string) {
	output := fmt.Sprintf("***** %s *****", val)
	line := strings.Repeat("-", len(output))

	fmt.Println(line)
	fmt.Println(output)
}
