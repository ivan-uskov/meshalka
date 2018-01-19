package actions

import (
	"net/http"
	"meshalka/model"
	"encoding/json"
	"fmt"
)

type addDrinkData struct {
	Name string `json:"name"`
	Type int `json:"type"`
}

func parseAddDrinkData(rawData string) (addDrinkData, error) {
	var data addDrinkData
	err := json.Unmarshal([]byte(rawData), &data)
	return data, err
}

func AddDrink(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	data, err := parseAddDrinkData(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	dr := model.NewDrinkRepository(rc.Ctx.Database)
	if !dr.IsNameValid(data.Name) {
		http.Error(writer, `{"status":"invalid_name"}`, http.StatusBadRequest)
		return
	}

	id, err := dr.Add(data.Name, data.Type)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	if id == 0 {
		http.Error(writer, `{"status":"name_busy"}`, http.StatusBadRequest)
		return
	}

	writer.Write([]byte(fmt.Sprintf(`{"status":"created","id":%d}`, id)))
}