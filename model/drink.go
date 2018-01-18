package model

import (
	"meshalka/database"
	"database/sql"
)

type Drink struct {
	Id uint64
	Name string
}

type DrinkRepository interface {
	Add(name string) (bool, error)
}

type drinkRepository struct {
	db database.Database
}

func NewDrinkRepository(db database.Database) DrinkRepository {
	return &drinkRepository{db}
}

func (dr * drinkRepository) Add(name string) (bool, error) {
	return getBoolResult(getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT add_drink(?)`, name)
	}))
}
