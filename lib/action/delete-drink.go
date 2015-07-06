package action

import (
	"fmt"
	"meshalka/lib/config"
	"meshalka/lib/database"
	"net/http"
)

func DeleteDrink(w http.ResponseWriter, r *http.Request, p config.Page) {
	drinkId := p.Args.Con["drink_id"]

	var result string
	err := database.RemoveDrink(drinkId)
	if err == nil {
		result = "{\"success\":1}"
	} else {
		result = "{\"success\":0}"
	}

	fmt.Fprintf(w, result)
}
