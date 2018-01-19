package actions

import (
	"net/http"
	"strconv"
	"meshalka/model"
	"fmt"
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
		http.NotFound(writer, request)
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"status":"deleted"}`)))
}