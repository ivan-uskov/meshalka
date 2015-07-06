package config

// mapped page id on javascript filenames
var scripts = map[string][]string{
	"index": []string{
		"lib/jquery-latest.min.js",
		"lib/jquery.fancybox.js",
		"lib/form.js",
		"lib/validators.js",
		"index.js",
	},
	"new-drink-form": []string{
		"new-drink.js",
	},
	"edit-drink-form": []string{
		"edit-drink.js",
	},
	"new-cocktail-form": []string{
		"new-cocktail.js",
	},
}

func GetPageScripts(pageId string) []string {
	return scripts[pageId]
}
