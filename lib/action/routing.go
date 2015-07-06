package action

import (
	"meshalka/lib/config"
	"meshalka/lib/utils"
)

var routes = []config.Route{
	config.Route{
		P: config.Page{
			Id:   "index",
			Path: "/",
			Args: config.UseArgs{
				Use: true,
				Con: make(map[string]string),
			},
		},
		H: Index,
	},
	config.Route{
		P: config.Page{
			Id:   "new-cocktail-form",
			Path: "/forms/new-cocktail",
			Args: config.UseArgs{
				Use: false,
				Con: make(map[string]string),
			},
		},
		H: NewCocktailForm,
	},
	config.Route{
		P: config.Page{
			Id:   "new-drink-form",
			Path: "/forms/new-drink",
			Args: config.UseArgs{
				Use: false,
				Con: make(map[string]string),
			},
		},
		H: NewDrinkForm,
	},
	config.Route{
		P: config.Page{
			Id:   "edit-drink-form",
			Path: "/forms/edit-drink/",
			Args: config.UseArgs{
				Use: true,
				Con: map[string]string{
					"drink_id": "[0-9]+",
				},
			},
		},
		H: EditDrinkForm,
	},
	config.Route{
		P: config.Page{
			Id:   "delete-drink",
			Path: "/forms/delete-drink/",
			Args: config.UseArgs{
				Use: true,
				Con: map[string]string{
					"drink_id": "[0-9]+",
				},
			},
		},
		H: DeleteDrink,
	},
	config.Route{
		P: config.Page{
			Id:   "new-drink-save",
			Path: "/ajax/save-drink",
			Args: config.UseArgs{
				Use: false,
				Con: make(map[string]string),
			},
		},
		H: NewDrinkSave,
	},
	config.Route{
		P: config.Page{
			Id:   "new-cocktail-save",
			Path: "/ajax/save-cocktail",
			Args: config.UseArgs{
				Use: false,
				Con: make(map[string]string),
			},
		},
		H: NewCocktailSave,
	},
	config.Route{
		P: config.Page{
			Id:   "edit-drink-save",
			Path: "/ajax/edit-drink/",
			Args: config.UseArgs{
				Use: true,
				Con: map[string]string{
					"drink_id": "[0-9]+",
				},
			},
		},
		H: EditDrinkSave,
	},
	config.Route{
		P: config.Page{
			Id:   "static",
			Path: "/static/",
			Args: config.UseArgs{
				Use: false,
				Con: make(map[string]string),
			},
		},
		H: Static,
	},
}

func AddHandlers() {
	for _, route := range routes {
		utils.AddHandler(route.P, route.H)
	}
}
