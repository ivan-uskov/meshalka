package config

// mapped page id on javascript filenames
var styles = map[string][]string{
	"index": []string{
		"index.css",
		"lib/jquery.fancybox.css",
	},
	"new-drink-form": []string{
		"new-drink.css",
	},
	"edit-drink-form": []string{
		"edit-drink.css",
	},
	"new-cocktail-form": []string{
		"new-cocktail.css",
	},
}

func GetPageStyles(pageId string) []string {
	return styles[pageId]
}
