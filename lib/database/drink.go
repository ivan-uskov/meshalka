package database

import (
	"database/sql"
	"fmt"
)

type Drink struct {
	Id   int
	Name string
}

func validateDrinkId(id string) error {
	return match(id, "^[0-9]+$", "Invalid drink id!")
}

func validateDrinkName(name string) error {
	return match(name, "^[0-9a-zA-Z-[[:space:]]+$", "Invalid drink name!")
}

func RemoveDrink(id string) error {
	con, err := getCon()
	if err != nil {
		return err
	}
	defer con.Close()

	return deleteDrink(id, con)
}

func deleteDrink(id string, con *sql.DB) error {
	err := validateDrinkId(id)
	if err != nil {
		return err
	}

	_, err = con.Exec(
		"DELETE FROM drink WHERE drink_id = (?)",
		id,
	)

	return err
}

func EditDrink(id string, newName string) error {
	con, err := getCon()
	if err != nil {
		return err
	}
	defer con.Close()

	return updateDrink(id, newName, con)
}

func updateDrink(id string, newName string, con *sql.DB) error {
	err := validateDrinkId(id)
	if err != nil {
		return err
	}

	err = validateDrinkName(newName)
	if err != nil {
		return err
	}

	_, err = con.Exec(
		"UPDATE drink SET drink_name = (?) WHERE drink_id = (?)",
		newName,
		id,
	)

	return err
}

func NewDrink(name string) error {
	con, err := getCon()
	if err != nil {
		return err
	}
	defer con.Close()

	return createDrink(name, con)
}

func createDrink(name string, con *sql.DB) error {
	err := validateDrinkName(name)
	if err != nil {
		return err
	}

	_, err = con.Exec(
		"INSERT INTO drink (drink_name) VALUES (?)",
		name,
	)

	return err
}

func GetDrinks() []Drink {
	con, err := getCon()
	if err != nil {
		return []Drink{}
	}
	defer con.Close()

	return selectDrinks(con)
}

func selectDrinks(con *sql.DB) []Drink {
	drinks := []Drink{}
	rows, err := con.Query("SELECT drink_id, drink_name FROM drink")
	if err != nil {
		fmt.Printf("Can't get drinks, ", err.Error())
		return drinks
	}
	defer rows.Close()

	drinks, err = fetchDrinks(rows)
	if err != nil {
		fmt.Printf("Can't parse drinks, ", err.Error())
	}

	return drinks
}

func fetchDrinks(rows *sql.Rows) ([]Drink, error) {
	drinks := []Drink{}
	var id int
	var name string
	var err error

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			break
		}
		drinks = append(drinks, Drink{
			Id:   id,
			Name: name,
		})
	}

	return drinks, err
}
