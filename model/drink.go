package model

import (
	"meshalka/database"
	"database/sql"
	"fmt"
)

const (
	maxNameLength = 255
	drinkType = 0
	cocktailType = 1
)

type Drink struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	Type int `json:"type"`
}

type DrinkRepository interface {
	IsNameValid(name string) bool
	Add(name string, t int) (int, error)
	List() (map[uint64]Drink, error)
	Remove(id int) (bool, error)
	Edit(id uint64, newName string) (int, error)
}

type drinkRepository struct {
	db database.Database
}

func NewDrinkRepository(db database.Database) DrinkRepository {
	return &drinkRepository{db}
}

func (dr * drinkRepository) IsNameValid(name string) bool {
	l := len(name)
	return l > 0  && l <= maxNameLength
}

func (dr * drinkRepository) Add(name string, t int) (int, error) {
	if !dr.IsNameValid(name) {
		return 0, fmt.Errorf("incorrect name")
	}

	if t != cocktailType {
		t = drinkType
	}

	res, err := getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT add_drink(?, ?)`, name, t)
	})

	if err != nil {
		return 0, err
	} else if res == -1 {
		return 0, fmt.Errorf("error while adding user with name %s and type %d", name, t)
	}

	return res, nil
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

func (dr * drinkRepository) Edit(id uint64, newName string) (int, error) {
	if !dr.IsNameValid(newName) {
		return 0, fmt.Errorf("incorrect new name")
	}

	res, err := getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT edit_drink(?, ?)`, id, newName)
	})

	if err != nil {
		return 0, err
	} else if res == -1 {
		return 0, fmt.Errorf("error while update user with new name %s and id %d", newName, id)
	}

	return res, nil
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