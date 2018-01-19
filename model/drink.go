package model

import (
	"meshalka/database"
	"database/sql"
	"fmt"
)

const (
	maxNameLength = 255
	drinkType     = 0
	cocktailType  = 1
)

type Drink struct {
	Id         uint64  `json:"id"`
	Name       string  `json:"name"`
	Type       int     `json:"type"`
	Components []Drink `json:"components"`
}

type DrinkRepository interface {
	IsNameValid(name string) bool
	Add(name string, t int) (int, error)
	List() (map[uint64]Drink, error)
	Remove(id int) (bool, error)
	Edit(id uint64, newName string) (int, error)
	AddCocktailElement(cocktailId uint64, drinkId uint64, quantity uint64) (int, error)
	RemoveWithExclusion(cocktailId uint64, excludedKeys string) error
}

type drinkRepository struct {
	db database.Database
}

func NewDrinkRepository(db database.Database) DrinkRepository {
	return &drinkRepository{db}
}

func (dr *drinkRepository) IsNameValid(name string) bool {
	l := len(name)
	return l > 0 && l <= maxNameLength
}

func (dr *drinkRepository) RemoveWithExclusion(cocktailId uint64, excludedKeys string) error {
	con, err := dr.db.Connection()
	if err != nil {
		return err
	}

	_, err = con.Query(`DELETE FROM cocktail_element WHERE cocktail_id = ? AND drink_id NOT IN (?)`, cocktailId, excludedKeys)
	return err
}

func (dr *drinkRepository) AddCocktailElement(cocktailId uint64, drinkId uint64, quantity uint64) (int, error) {
	return getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT add_cocktail_element(?, ?, ?)`, cocktailId, drinkId, quantity)
	})
}

func (dr *drinkRepository) Add(name string, t int) (int, error) {
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

func (dr *drinkRepository) Remove(id int) (bool, error) {
	return getBoolResult(getIntFunctionResult(dr.db, func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`SELECT remove_drink(?)`, id)
	}))
}

func (dr *drinkRepository) List() (map[uint64]Drink, error) {
	return dr.selectDrinks(func(con *sql.DB) (*sql.Rows, error) {
		return con.Query(`
			SELECT
			  d.drink_id,
			  d.drink_name,
			  d.drink_type,
			  COALESCE(subD.drink_id, 0),
			  COALESCE(subD.drink_name, '')
			FROM
			  drink d
			  LEFT JOIN cocktail_element ce ON d.drink_id = ce.cocktail_id
			  LEFT JOIN drink subD ON ce.drink_id = subD.drink_id
			ORDER BY
			  d.drink_type ASC
		`)
	})
}

func (dr *drinkRepository) Edit(id uint64, newName string) (int, error) {
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
	drink := Drink{Components:[]Drink{}}
	var subDrink Drink
	drinks := make(map[uint64]Drink)

	for rows.Next() {
		err := rows.Scan(&drink.Id, &drink.Name, &drink.Type, &subDrink.Id, &subDrink.Name)
		if err != nil {
			return nil, err
		}

		if drink.Type == cocktailType && subDrink.Id > 0 {
			drink.Components = append(drink.Components, subDrink)
		}
		drinks[drink.Id] = drink
	}

	return drinks, nil
}
