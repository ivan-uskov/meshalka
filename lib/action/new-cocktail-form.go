package action

import (
	"meshalka/lib/config"
	"meshalka/lib/database"
	"meshalka/lib/utils/templates"
	"net/http"
)

func NewCocktailForm(w http.ResponseWriter, r *http.Request, p config.Page) {
	fi := config.NewFormInfo(p.Id, "new-cocktail.html", database.GetDrinks())

	templates.RenderForm(w, fi)
}
