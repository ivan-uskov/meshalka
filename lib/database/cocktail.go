package database

import (
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"os/exec"
)

type Cocktail struct {
	Id   int
	Name string
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
	var out bytes.Buffer

	cmd := exec.Command("/usr/bin/php", "./add_cocktail.php", name, drinks)
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	return parseNewCocktailOutput(out.String())
}

func parseNewCocktailOutput(out string) error {
	fmt.Println("Add cocktail reported: " + out)
	if out != "1" {
		return errors.New("Cocktail creation error!")
	}

	return nil
}

func GetCocktail() []Cocktail {
	con, err := getCon()
	if err != nil {
		return []Cocktail{}
	}
	defer con.Close()

	return selectCocktail(con)
}

func selectCocktail(con *sql.DB) []Cocktail {
	cocktails := []Cocktail{}
	rows, err := con.Query("SELECT cocktail_id, cocktail_name FROM cocktail")
	if err != nil {
		fmt.Printf("Can't get cocktails, ", err.Error())
		return cocktails
	}
	defer rows.Close()

	cocktails, err = fetchCocktail(rows)
	if err != nil {
		fmt.Printf("Can't parse cocktails, ", err.Error())
	}

	return cocktails
}

func fetchCocktail(rows *sql.Rows) ([]Cocktail, error) {
	cocktails := []Cocktail{}
	var id int
	var name string
	var err error

	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			break
		}
		cocktails = append(cocktails, Cocktail{
			Id:   id,
			Name: name,
		})
	}

	return cocktails, err
}
