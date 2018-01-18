package action

import (
	"net/http"
	"encoding/json"
	"meshalka/model"
	"fmt"
)

type loginData struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
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
		writer.Write([]byte(fmt.Sprintf(`{"login":"%s","user_id":"%d"}`, user.Login, user.UserId)))
	} else {
		http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
	}
}