package templates

import (
	"meshalka/lib/config"
	"net/http"
)

func RenderPage(w http.ResponseWriter, pageInfo config.PageInfo) {
	tmpl, ok := templates[pageInfo.TplName]

	if ok {
		tmpl.ExecuteTemplate(w, "base", pageInfo)
	} else {
		http.Error(w, "The template "+pageInfo.TplName+"does not exist.", http.StatusInternalServerError)
	}
}

func RenderForm(w http.ResponseWriter, formInfo config.FormInfo) {
	tmpl, ok := templates["form-"+formInfo.TplName]

	if ok {
		tmpl.ExecuteTemplate(w, "base", formInfo)
	} else {
		http.Error(w, "The template "+formInfo.TplName+"does not exist.", http.StatusInternalServerError)
	}
}
