package main

import (
	"log"
	"net/http"

	"github.com/go-kit/kit/metrics/provider"
	"github.com/tomdevito/sizzle/api"
	"github.com/tomdevito/sizzle/app"
	"github.com/tomdevito/sizzle/store"
)

func main() {
	//configure

	// create app
	app := &app.App{
		Store:  &store.SimpleStore{},
		Metric: provider.NewDiscardProvider(),
	}

	api := api.NewAPI(app)
	apiHandler := api.Initialize()
	// api will contain complete handler to mount onto

	log.Fatal(http.ListenAndServe(":8080", apiHandler))
}
