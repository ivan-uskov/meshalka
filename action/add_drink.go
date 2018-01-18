package action

import (
	"net/http"
	"meshalka/model"
)

func AddDrink(writer http.ResponseWriter, rc *RequestContext) {
	added, err := model.NewDrinkRepository(rc.Ctx.Database).Add(rc.Data)
	if err != nil {
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	if !added {
		http.Error(writer, "Drink already exists", http.StatusBadRequest)
	}
}