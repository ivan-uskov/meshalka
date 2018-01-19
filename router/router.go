package router

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"meshalka/contexts"
	"meshalka/actions"
)

const (
	loginAction                  = `login`
	registerAction               = `register`
	addDrinkAction               = `add_drink`
	removeDrinkAction            = `remove_drink`
	editDrinkAction              = `edit_drink`
	editCocktailComponentsAction = `edit_cocktail_components`
	listAction                   = `list`
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
	return !(action == loginAction || action == registerAction)
}

func (r *router) executeAction(w http.ResponseWriter, hr *http.Request, req *requestBody) {
	rc := &actions.RequestContext{Data: req.Data}
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
		actions.Login(w, hr, rc)
	case registerAction:
		actions.Registration(w, hr, rc)
	case addDrinkAction:
		actions.AddDrink(w, hr, rc)
	case listAction:
		actions.List(w, rc)
	case removeDrinkAction:
		actions.RemoveDrink(w, hr, rc)
	case editDrinkAction:
		actions.EditDrink(w, hr, rc)
	case editCocktailComponentsAction:
		actions.EditCocktailComponents(w, hr, rc)
	default:
		http.NotFound(w, hr)
	}
}
