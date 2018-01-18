package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"meshalka/config"
)

type Database interface {
	Connection() (*sql.DB, error)
}

func New(conf *config.DBConfig) (Database, error) {
	db, err := createConnection(conf)
	if err != nil {
		return nil, err
	}

	return &database{db, conf}, nil
}

type database struct {
	pool *sql.DB
	conf *config.DBConfig
}

func (db *database) Connection() (*sql.DB, error) {
	err := db.pool.Ping()
	if err != nil {
		db.pool.Close()
		db.pool, err = createConnection(db.conf)
	}

	return db.pool, err
}

func createConnection(conf *config.DBConfig) (*sql.DB, error) {
	con, err := sql.Open(conf.Driver, conf.Dsn)
	if err != nil {
		return nil, err
	}

	return con, con.Ping()
}