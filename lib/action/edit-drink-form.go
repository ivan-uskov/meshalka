package action

import (
	"meshalka/lib/config"
	"meshalka/lib/utils/templates"
	"net/http"
)

type String struct {
	Val string
}

func EditDrinkForm(w http.ResponseWriter, r *http.Request, p config.Page) {
	fi := config.NewFormInfo(p.Id, "edit-drink.html", String{
		Val: p.Args.Con["drink_id"],
	})

	templates.RenderForm(w, fi)
}
