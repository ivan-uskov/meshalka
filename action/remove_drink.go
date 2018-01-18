package action

import (
	"net/http"
	"strconv"
	"meshalka/model"
)

func RemoveDrink(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	id, err := strconv.Atoi(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	removed, err := model.NewDrinkRepository(rc.Ctx.Database).Remove(id)
	if err != nil {
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	if !removed {
		http.Error(writer, "Drink not exists", http.StatusNoContent)
	}
}