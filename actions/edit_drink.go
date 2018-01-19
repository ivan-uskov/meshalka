package actions

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
	"fmt"
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

	dr := model.NewDrinkRepository(rc.Ctx.Database)
	if !dr.IsNameValid(data.NewName) {
		http.Error(writer, `{"status":"invalid_name"}`, http.StatusBadRequest)
		return
	}

	res, err := model.NewDrinkRepository(rc.Ctx.Database).Edit(data.Id, data.NewName)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}
	if res == -2 { //drink not exist
		http.NotFound(writer, request)
		return
	}

	if res == 0 {
		http.Error(writer, `{"status":"name_busy"}`, http.StatusBadRequest)
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"status":"updated"}`)))
}
