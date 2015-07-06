package action

import (
	"fmt"
	"meshalka/lib/config"
	"meshalka/lib/database"
	"net/http"
)

func NewCocktailSave(w http.ResponseWriter, r *http.Request, p config.Page) {
	cocktailName := r.FormValue("cocktail_name")
	drinks := r.FormValue("drinks")

	var result string
	err := database.NewCocktail(cocktailName, drinks)
	if err == nil {
		result = "{\"success\":1}"
	} else {
		fmt.Println(err.Error())
		result = "{\"success\":0}"
	}

	fmt.Fprintf(w, result)
}
