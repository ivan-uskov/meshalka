package action

import (
	"fmt"
	"meshalka/lib/config"
	"meshalka/lib/database"
	"net/http"
)

func NewDrinkSave(w http.ResponseWriter, r *http.Request, p config.Page) {
	drinkName := r.FormValue("drink_name")

	var result string
	err := database.NewDrink(drinkName)
	if err == nil {
		result = "{\"success\":1}"
	} else {
		result = "{\"success\":0}"
	}

	fmt.Fprintf(w, result)
}
