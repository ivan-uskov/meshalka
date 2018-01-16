package action

import (
	"meshalka/lib/config"
	"meshalka/lib/utils/templates"
	"net/http"
	"meshalka/lib/database"
	"strconv"
)

type Data struct {
	Drink database.Drink
}

func EditDrinkForm(w http.ResponseWriter, r *http.Request, p config.Page) {
	id, err := strconv.Atoi(p.Args.Con["drink_id"])
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	drink, err := database.GetDrink(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	templates.RenderForm(w, config.NewFormInfo(p.Id, "edit-drink.html", Data{*drink}))
}
