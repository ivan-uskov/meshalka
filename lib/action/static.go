package action

import (
	"meshalka/lib/config"
	"net/http"
)

func Static(w http.ResponseWriter, r *http.Request, p config.Page) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
