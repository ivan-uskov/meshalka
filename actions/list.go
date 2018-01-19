package actions

import (
	"net/http"
	"meshalka/model"
	"encoding/json"
	"fmt"
)

type list struct {
	Drinks []model.Drink `json:"drinks"`
}

func prepareResponse(drinks map[uint64]model.Drink) ([]byte, error) {
	drinksList := make([]model.Drink, len(drinks))
	idx := 0
	for  _, value := range drinks {
		drinksList[idx] = value
		idx++
	}

	return json.Marshal(list{drinksList})
}

func List(writer http.ResponseWriter, rc *RequestContext) {
	drinks, err := model.NewDrinkRepository(rc.Ctx.Database).List()
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	resp, err := prepareResponse(drinks)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(resp)
}
