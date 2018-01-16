package action

import (
	"meshalka/lib/config"
	"meshalka/lib/utils/templates"
	"net/http"
)

func NewDrinkForm(w http.ResponseWriter, r *http.Request, p config.Page) {
	fi := config.NewFormInfo(p.Id, "new-drink.html", struct{}{})

	templates.RenderForm(w, fi)
}
