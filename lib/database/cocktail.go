package database

import (
	"database/sql"
	"fmt"
)

type Cocktail struct {
	Id   int
	Name string
	Drinks map[int]Drink
}

func validateCocktailId(id string) error {
	return match(id, "^[0-9]+$", "Invalid drink id!")
}

func validateCoctailName(name string) error {
	return match(name, "^[0-9a-zA-Z-[[:space:]]+$", "Invalid drink name!")
}

func validateDrinks(drinks string) error {
	return match(drinks, "^[0-9]([,][0-9])*$", "Invalid drinks string!")
}

func NewCocktail(name string, drinks string) error {
	con, err := getCon()
	if err != nil {
		return err
	}
	defer con.Close()

	_, err = con.Exec("CALL add_cocktail(?, ?)", name, drinks)

	return err
}

func GetCocktails() map[int]Cocktail {
	con, err := getCon()
	if err != nil {
		return make(map[int]Cocktail)
	}
	defer con.Close()

	return selectCocktails(con)
}

func selectCocktails(con *sql.DB) map[int]Cocktail {
	rows, err := con.Query("SELECT cocktail_id, cocktail_name, d.drink_id, d.drink_name FROM cocktail LEFT JOIN cocktail_element USING(cocktail_id) LEFT JOIN drink d USING(drink_id)")
	if err != nil {
		fmt.Printf("Can't get cocktails, %s", err.Error())
		return make(map[int]Cocktail)
	}
	defer rows.Close()

	cocktails, err := fetchCocktails(rows)
	if err != nil {
		fmt.Printf("Can't parse cocktails, %s", err.Error())
	}

	return cocktails
}

func fetchCocktails(rows *sql.Rows) (map[int]Cocktail, error) {
	cocktails := make(map[int]Cocktail)
	var id int
	var name string
	var drinkId int
	var drinkName string
	var err error

	for rows.Next() {
		err = rows.Scan(&id, &name, &drinkId, &drinkName)
		if err != nil {
			break
		}

		_, ok := cocktails[id]
		if !ok {
			cocktails[id] = Cocktail{id, name, make(map[int]Drink)}
		}
		cocktails[id].Drinks[drinkId] = Drink{drinkId, drinkName}
	}

	return cocktails, err
}
