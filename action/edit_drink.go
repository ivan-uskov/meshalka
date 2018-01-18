package action

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
)

type editDrinkData struct {
	Id uint64 `json:"id"`
	NewName string `json:"new_name"`
}

func parseEditDrinkData(rawData string) (editDrinkData, error) {
	var data editDrinkData
	err := json.Unmarshal([]byte(rawData), &data)
	return data, err
}

func EditDrink(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	data, err := parseEditDrinkData(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	added, err := model.NewDrinkRepository(rc.Ctx.Database).Edit(data.Id, data.NewName)
	if err != nil {
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	if !added {
		http.Error(writer, "Name busy or drink not exists", http.StatusBadRequest)
	}
}
