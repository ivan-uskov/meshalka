package actions

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
	"fmt"
	"strings"
)

type editCocktailComponentsData struct {
	Id uint64 `json:"id"`
	Components map[uint64]uint64 `json:"components"`
}

func parseEditCocktailComponentsData(rawData string) (*editCocktailComponentsData, error) {
	data := &editCocktailComponentsData{Components: make(map[uint64]uint64)}
	err := json.Unmarshal([]byte(rawData), data)
	return data, err
}

func EditCocktailComponents(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	data, err := parseEditCocktailComponentsData(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	dr := model.NewDrinkRepository(rc.Ctx.Database)

	i := 0
	keys := make([]string, len(data.Components))
	for drinkId, quantity := range data.Components {
		res, err := dr.AddCocktailElement(data.Id, drinkId, quantity)
		if err != nil || res == -1{
			fmt.Println(err)
			http.Error(writer, "Internal error", http.StatusInternalServerError)
			return
		} else if res == -2 {
			http.Error(writer, `{"status":"cocktail_not_exists"}`, http.StatusBadRequest)
			return
		} else if res == -3 {
			http.Error(writer, `{"status":"is_not_a_cocktail"}`, http.StatusBadRequest)
			return
		} else if res == 1 {
			keys[i] = fmt.Sprintf("%d", drinkId)
			i++
		}
	}

	excludedKeys := strings.Join(keys[:i], ",")
	err = dr.RemoveWithExclusion(data.Id, excludedKeys)
	if err != nil {
		fmt.Println(err)
	}

	writer.Write([]byte(fmt.Sprintf(`{"status":"updated"}`)))
}