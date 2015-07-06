package action

import (
	"fmt"
	"meshalka/lib/config"
	"meshalka/lib/database"
	"net/http"
)

func EditDrinkSave(w http.ResponseWriter, r *http.Request, p config.Page) {
	drinkName := r.FormValue("drink_name")
	drinkId := p.Args.Con["drink_id"]

	var result string
	err := database.EditDrink(drinkId, drinkName)
	if err == nil {
		result = "{\"success\":1}"
	} else {
		result = "{\"success\":0}"
	}

	fmt.Fprintf(w, result)
}
