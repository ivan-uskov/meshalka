package templates

import (
	"html/template"
	"path/filepath"
)

var templates map[string]*template.Template

func loadDir(keyPrefix string, baseDir string) {

	baseTplPath := baseDir + "base.tpl"

	pages, err := filepath.Glob(baseDir + "*.html")
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		templates[keyPrefix+filepath.Base(page)] = template.Must(template.ParseFiles(page, baseTplPath))
	}
}

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templateDir := "static/templates/"

	loadDir("", templateDir)
	loadDir("form-", templateDir+"forms/")
}
