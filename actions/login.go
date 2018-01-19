package actions

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
)

type loginData struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type loginResponse struct {
	Login string `json:"login"`
	UserId uint64 `json:"user_id"`
}

func parseLoginData(data string) (*loginData, error) {
	ld := &loginData{}
	err := json.Unmarshal([]byte(data), ld)
	return ld, err
}

func Login(writer http.ResponseWriter, request *http.Request, rc *RequestContext) {
	ld, err := parseLoginData(rc.Data)
	if err != nil {
		http.NotFound(writer, request)
		return
	}

	user, err := model.NewUserRepository(rc.Ctx.Database).SelectUserByLoginInfo(ld.Login, ld.Pass)
	if err != nil {
		http.Error(writer, "Not authorized", http.StatusUnauthorized)
		return
	}

	if rc.Ctx.Session.Login(writer, request, user) {
		data, err := json.Marshal(loginResponse{user.Login, user.UserId})
		if err != nil {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		writer.Write(data)
	} else {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}