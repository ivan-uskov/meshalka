package main

import (
	"net/http"
	"fmt"
	"os"
	"meshalka/contexts"
	gorillaContext "github.com/gorilla/context"
	"log"
	"meshalka/router"
)

func main() {
	ctx, err := contexts.NewContext(`config.yml`)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	router.New(ctx).Register()
	log.Fatal(http.ListenAndServe(ctx.Config.AppAddress, gorillaContext.ClearHandler(http.DefaultServeMux)))
}
