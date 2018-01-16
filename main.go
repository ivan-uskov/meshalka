package main

import (
	"meshalka/lib/action"
	"net/http"
)

func main() {
	action.AddHandlers()
	http.ListenAndServe(":8080", nil)
}
