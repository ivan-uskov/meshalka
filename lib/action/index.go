package action

import (
	"meshalka/lib/config"
	"meshalka/lib/database"
	"meshalka/lib/utils/templates"
	"net/http"
)

type indexData struct {
	Drinks    []database.Drink
	Cocktails []database.Cocktail
}

func Index(w http.ResponseWriter, r *http.Request, p config.Page) {
	data := indexData{
		Drinks:    database.GetDrinks(),
		Cocktails: database.GetCocktail(),
	}

	pi := config.NewPageInfo(p.Id, "Cocktails", "index.html", data)

	templates.RenderPage(w, pi)
}
