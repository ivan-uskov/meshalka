package router

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"meshalka/contexts"
	"meshalka/action"
)

const (
	loginAction = `login`
	registerAction = `register`
	addDrinkAction = `add_drink`
	listAction = `list`
)

type requestBody struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

type router struct {
	ctx *contexts.Context
}

type Router interface {
	Register()
}

func New(ctx *contexts.Context) Router {
	return &router{ctx}
}

func (r *router) Register() {
	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.NotFound(writer, request)
			return
		}

		reqBody := &requestBody{}
		err = json.Unmarshal(body, reqBody)
		if err != nil {
			http.NotFound(writer, request)
			return
		}

		r.executeAction(writer, request, reqBody)
	})
}

func isPrivateAction(action string) bool {
	return false
}

func (r *router) executeAction(w http.ResponseWriter, hr *http.Request, req *requestBody) {
	rc := &action.RequestContext{Data: req.Data}
	if isPrivateAction(req.Action) {
		if user, ok := r.ctx.Session.AutoLoginFilter(hr); ok {
			rc.User = user
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	}

	rc.Ctx = r.ctx

	switch req.Action {
	case loginAction:
		action.Login(w, hr, rc)
	case registerAction:
		action.Registration(w, hr, rc)
	case addDrinkAction:
		action.AddDrink(w, hr, rc)
	case listAction:
		action.List(w, hr, rc)
	default:
		http.NotFound(w, hr)
	}
}
