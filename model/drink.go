package model

import (
	"meshalka/database"
	"database/sql"
	"fmt"
)

const maxNameLength = 255

type Drink struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
}

type DrinkRepository interface {
	Add(name string) (bool, error)
	List() (map[uint64]Drink, error)
	Remove(id int) (bool, error)
	Edit(id uint64, newName string) (bool, error)
}

type drinkRepository struct {
	db database.Database
}

func NewDrinkRepository(db database.Database) DrinkRepository {
	return &drinkRepository{db}
}

func (dr * drinkRepository) Add(name string) (bool, error) {
	if len(name) >= maxNameLength {
		return false, fmt.Errorf("name too long")
	}

	return getBoolResult(getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT add_drink(?)`, name)
	}))
}

func (dr * drinkRepository) Remove(id int) (bool, error) {
	return getBoolResult(getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT remove_drink(?)`, id)
	}))
}

func (dr * drinkRepository) List() (map[uint64]Drink, error) {
	return dr.selectDrinks(func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT drink_id, drink_name FROM drink`)
	})
}

func (dr * drinkRepository) Edit(id uint64, newName string) (bool, error) {
	if len(newName) >= maxNameLength {
		return false, fmt.Errorf("name too long")
	}

	return getBoolResult(getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT edit_drink(?, ?)`, id, newName)
	}))
}

func (dr *drinkRepository) selectDrinks(q querier) (map[uint64]Drink, error) {
	con, err := dr.db.Connection()
	if err != nil {
		return nil, err
	}

	rows, err := q(con)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return fetchDrinks(rows)
}

func fetchDrinks(rows *sql.Rows) (map[uint64]Drink, error) {
	var drink Drink
	drinks := make(map[uint64]Drink)

	for rows.Next() {
		err := rows.Scan(&drink.Id, &drink.Name)
		if err != nil {
			return nil, err
		}

		drinks[drink.Id] = drink
	}

	return drinks, nil
}