package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-kit/kit/metrics/provider"
	"github.com/tomdevito/sizzle/api"
	"github.com/tomdevito/sizzle/app"
	"github.com/tomdevito/sizzle/store"
)

func main() {
	//configure / create app
	app := &app.App{
		Store:  &store.SimpleStore{},
		Metric: provider.NewDiscardProvider(),
	}

	api := api.NewAPI(app)

	// api will contain complete handler to mount onto
	r := chi.NewRouter()
	r.Mount("/smoke", api.LoadSmokeHandler())

	log.Fatal(http.ListenAndServe(":8080", r))
}
