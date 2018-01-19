package actions

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
)

type registrationData struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func parseRegistrationData(data string) (*registrationData, error) {
	rd := &registrationData{}
	err := json.Unmarshal([]byte(data), rd)
	return rd, err
}

func Registration(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	rd, err := parseRegistrationData(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	added, err := model.NewUserRepository(rc.Ctx.Database).AddUser(rd.Login, rd.Pass)
	if err != nil {
		http.Error(writer, "Internal error", http.StatusInternalServerError)
		return
	}

	if !added {
		http.Error(writer, "User already exists", http.StatusBadRequest)
	}
}