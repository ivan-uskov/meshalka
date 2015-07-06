package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"regexp"
)

func match(str string, pattern string, msg string) error {
	if regexp.MustCompile(pattern).MatchString(str) {
		return nil
	}

	return errors.New(msg)
}

func getCon() (*sql.DB, error) {
	con, err := sql.Open("mysql", "root:1234@/cocktails")
	if err != nil {
		fmt.Println("Can't init database connection")
		return nil, err
	}

	return con, nil
}

func init() {
	con, err := getCon()
	if err != nil {
		panic(err.Error())
	}

	err = con.Ping()
	if err != nil {
		panic(err.Error())
	}
}
